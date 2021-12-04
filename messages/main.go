package main

import (
	"github.com/Reoneks/microservice_chat/messages/config"
	"github.com/Reoneks/microservice_chat/messages/messages"
	"github.com/Reoneks/microservice_chat/proto"
	"github.com/asim/go-micro/v3"
)

func main() {
	cfg := config.GetConfig()
	db, err := config.NewDB(cfg.DSN, cfg.MigrationURL)
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

	rabbitMicro := messages.NewMessagesMicro(&cfg, msgService, amqpChan, &queue, nil)
	microService := micro.NewService(micro.Name(cfg.ServiceName))
	microService.Init()

	if err := proto.RegisterMessagesHandler(
		microService.Server(),
		&messages.MessagesMicro{MessagesService: msgService},
	); err != nil {
		panic(err)
	}

	go func() {
		if err := rabbitMicro.StartConsumer(); err != nil {
			panic(err)
		}
	}()

	if err := microService.Run(); err != nil {
		panic(err)
	}
}
