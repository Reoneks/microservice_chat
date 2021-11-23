package messages

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Reoneks/microservice_chat/messages/config"
	"github.com/Reoneks/microservice_chat/messages/model"
	"github.com/Reoneks/microservice_chat/proto"

	"github.com/streadway/amqp"
)

type messagesMicro struct {
	MessagesService MessagesService

	Ch *amqp.Channel
	Q  *amqp.Queue

	//& StartConsumer
	Consumer  string     //^ ""
	AutoAsk   bool       //^ true
	Exclusive bool       //^ false
	NoLocal   bool       //^ false
	NoWait    bool       //^ false
	Arguments amqp.Table //^ nil

	//& publish
	Exchange  string //^ ""
	Mandatory bool   //^ false
	Immediate bool   //^ false
}

type IMessagesMicro interface {
	StartConsumer() error
}

func NewMessagesMicro(
	cfg *config.Config,
	MessagesService MessagesService,
	Ch *amqp.Channel,
	Q *amqp.Queue,
	Arguments amqp.Table,
) IMessagesMicro {
	return &messagesMicro{
		MessagesService: MessagesService,

		Ch: Ch,
		Q:  Q,

		Consumer:  cfg.Consumer,
		AutoAsk:   cfg.AutoAsk,
		Exclusive: cfg.ConsumerExclusive,
		NoLocal:   cfg.NoLocal,
		NoWait:    cfg.ConsumerNoWait,
		Arguments: Arguments,

		Exchange:  cfg.Exchange,
		Mandatory: cfg.Mandatory,
		Immediate: cfg.Immediate,
	}
}

func (u *messagesMicro) publish(data []byte) error {
	return u.Ch.Publish(
		u.Exchange,
		"",
		u.Mandatory,
		u.Immediate,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		},
	)
}

func (u *messagesMicro) createMessage(data *model.Message) error {
	message, err := u.MessagesService.CreateMessage(data)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return u.publish(bytes)
}

func (u *messagesMicro) updateMessage(data *model.Message) error {
	message, err := u.MessagesService.UpdateMessage(data)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return u.publish(bytes)
}

func (u *messagesMicro) deleteMessage(id string) error {
	err := u.MessagesService.DeleteMessage(id)
	if err != nil {
		return err
	}

	return nil
}

func (u *messagesMicro) StartConsumer() error {
	msgs, err := u.Ch.Consume(
		u.Q.Name,
		u.Consumer,
		u.AutoAsk,
		u.Exclusive,
		u.NoLocal,
		u.NoWait,
		u.Arguments,
	)
	if err != nil {
		return err
	}

	for message := range msgs {
		var data model.Message

		err := json.Unmarshal(message.Body, &data)
		if err == nil {
			switch data.MessageType {

			case "create":
				err := u.createMessage(&data)
				if err != nil {
					log.Printf("%v", err)
				}

			case "update":
				err := u.updateMessage(&data)
				if err != nil {
					log.Printf("%v", err)
				}

			case "delete":
				err := u.deleteMessage(data.ID)
				if err != nil {
					log.Printf("%v", err)
				}

			}
		}
	}

	return nil
}

type MessagesMicro struct {
	MessagesService MessagesService
}

func (u *MessagesMicro) GetMessagesByRoom(
	ctx context.Context,
	req *proto.MessageID,
	rsp *proto.GetMessagesByRoomResponse,
) error {
	messages, err := u.MessagesService.GetMessagesByRoom(req.ID, int(req.Limit), int((req.Page-1)*req.Limit))
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	bytes, err := json.Marshal(messages)
	if err != nil {
		rsp.Status.Ok = false
		rsp.Status.Error = err.Error()
		return err
	}

	rsp.Messages = bytes
	rsp.Status.Ok = true
	return nil
}
