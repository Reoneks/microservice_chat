package messages

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Reoneks/microservice_chat/messages/config"
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

func (u *messagesMicro) publish(data []byte, msg *proto.RabbitMessage) error {
	msg.Message = data

	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return u.Ch.Publish(
		u.Exchange,
		"",
		u.Mandatory,
		u.Immediate,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bytes,
		},
	)
}

func (u *messagesMicro) createMessage(data *proto.RabbitMessage) error {
	var msg map[string]interface{}
	err := json.Unmarshal(data.Message, &msg)
	if err != nil {
		log.Println("Unmarshal msg error: ", err)
	}

	message, err := u.MessagesService.CreateMessage(msg)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return u.publish(bytes, data)
}

func (u *messagesMicro) updateMessage(data *proto.RabbitMessage) error {
	var msg map[string]interface{}
	err := json.Unmarshal(data.Message, &msg)
	if err != nil {
		log.Println("Unmarshal msg error: ", err)
	}

	message, err := u.MessagesService.UpdateMessage(msg)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return u.publish(bytes, data)
}

func (u *messagesMicro) deleteMessage(data *proto.RabbitMessage) error {
	err := u.MessagesService.DeleteMessage(string(data.Message))
	if err != nil {
		return err
	}

	return u.publish(nil, data)
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
		var data proto.RabbitMessage

		err := json.Unmarshal(message.Body, &data)
		if err != nil {
			log.Println("Unmarshal RabbitMessage error: ", err)
		}

		switch data.MessageType {
		case "create":
			var msg map[string]interface{}
			err = json.Unmarshal(data.Message, &msg)
			if err != nil {
				log.Println("Unmarshal msg error: ", err)
			}

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
			err := u.deleteMessage(&data)
			if err != nil {
				log.Printf("%v", err)
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
	rsp.Status = new(proto.Status)

	messages, err := u.MessagesService.GetMessagesByRoom(req.RoomID, req.Limit, req.Offset)
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
