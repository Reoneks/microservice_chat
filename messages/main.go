package main

import (
	"github.com/Reoneks/microservice_chat/messages/config"
	"github.com/Reoneks/microservice_chat/messages/messages"
)

func main() {
	cfg := config.GetConfig()
	db, err := config.NewDB(cfg.DSN)
	if err != nil {
		panic(err)
	}

	amqpChan, err := config.StartRabbitMQ(cfg.RabbitMQUrl)
	if err != nil {
		panic(err)
	}

	msgRep := messages.NewMessagesRepository(db)
	msgService := messages.NewMessagesService(msgRep)

	amqp := config.GetAMQP(&cfg, nil, nil)
	queue, err := amqp.GetQueue(amqpChan)
	if err != nil {
		panic(err)
	}

	if err := amqp.SetExchange(amqpChan); err != nil {
		panic(err)
	}

	micro := messages.NewMessagesMicro(&cfg, msgService, amqpChan, &queue, nil)

	if err := micro.StartConsumer(); err != nil {
		panic(err)
	}
}
