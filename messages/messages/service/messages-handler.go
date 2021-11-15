package service

import (
	"chatex/messages/messages/messages_interface"
	"encoding/json"
	"log"

	"github.com/mitchellh/mapstructure"
	"github.com/streadway/amqp"
)

type MessagesMicro struct {
	MessagesService messages_interface.MessagesService

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

func (u *MessagesMicro) publish(msg amqp.Delivery, data []byte) error {
	return u.Ch.Publish(
		u.Exchange,
		msg.ReplyTo,
		u.Mandatory,
		u.Immediate,
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: msg.CorrelationId,
			Body:          data,
		},
	)
}

func (u *MessagesMicro) createMessage(data map[string]interface{}, msg amqp.Delivery) error {
	var message *messages_interface.Message

	err := mapstructure.Decode(data, message)
	if err != nil {
		return err
	}

	message, err = u.MessagesService.CreateMessage(message)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return u.publish(msg, bytes)
}

func (u *MessagesMicro) updateMessage(data map[string]interface{}, msg amqp.Delivery) error {
	var message *messages_interface.Message

	err := mapstructure.Decode(data, message)
	if err != nil {
		return err
	}

	message, err = u.MessagesService.UpdateMessage(message)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return u.publish(msg, bytes)
}

func (u *MessagesMicro) deleteMessage(id string, msg amqp.Delivery) error {
	err := u.MessagesService.DeleteMessage(id)
	if err != nil {
		return err
	}

	return u.publish(msg, nil)
}

func (u *MessagesMicro) StartConsumer() error {
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
		var data map[string]interface{}

		err := json.Unmarshal(message.Body, &data)
		if err == nil {
			switch data["message_type"].(string) {

			case "create":
				err := u.createMessage(data, message)
				if err != nil {
					log.Printf("%v", err)
				}

			case "update":
				err := u.updateMessage(data, message)
				if err != nil {
					log.Printf("%v", err)
				}

			case "delete":
				err := u.deleteMessage(data["id"].(string), message)
				if err != nil {
					log.Printf("%v", err)
				}

			}
		}
	}

	return nil
}
