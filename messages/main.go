package main

import (
	"chatex/messages/config"
	"chatex/messages/messages/repository"
	"chatex/messages/messages/service"
)

func main() {
	cfg := config.GetConfig()
	db, err := config.NewDB(cfg.DSN)
	if err != nil {
		panic(err)
	}

	rabbitmq, err := config.StartRabbitMQ(cfg.RabbitMQUrl)
	if err != nil {
		panic(err)
	}

	msgRep := repository.NewMessagesRepository(db)
	msgService := service.NewMessagesService(msgRep)

	amqp := config.AmqpSettings{
		Name:          "messages",
		Exclusive:     true,
		PrefetchCount: 1,
	}
	queue, err := amqp.GetQueue(rabbitmq)
	if err != nil {
		panic(err)
	}

	micro := service.MessagesMicro{
		MessagesService: msgService,
		Ch:              rabbitmq,
		Q:               &queue,
		AutoAsk:         true,
	}

	if err := micro.StartConsumer(); err != nil {
		panic(err)
	}
}
