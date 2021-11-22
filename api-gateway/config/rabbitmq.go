package config

import "github.com/streadway/amqp"

func StartRabbitMQ(url string) (*amqp.Channel, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	return conn.Channel()
}

type amqpSettings struct {
	//& Send Queue
	SendName             string     //^ messages
	SendDurable          bool       //^ false
	SendDeleteWhenUnused bool       //^ false
	SendExclusive        bool       //^ true
	SendNoWait           bool       //^ false
	QSendArguments       amqp.Table //^ nil

	//& Receive Queue
	ReceiveName             string     //^ messages
	ReceiveDurable          bool       //^ false
	ReceiveDeleteWhenUnused bool       //^ false
	ReceiveExclusive        bool       //^ true
	ReceiveNoWait           bool       //^ false
	QReceiveArguments       amqp.Table //^ nil

	//& ExchangeDeclare
	ExchangeName      string //^ messages-gateway
	Type              string //^ fanout
	ExchangeDurable   bool   //^ true
	AutoDelete        bool   //^ false
	Internal          bool   //^ false
	ExchangeNoWait    bool   //^ false
	ExchangeArguments amqp.Table
}

type IAMQP interface {
	GetSendQueue(ch *amqp.Channel) (amqp.Queue, error)
	GetReceiveChan(ch *amqp.Channel) (<-chan amqp.Delivery, error)
}

func GetAMQP(cfg *Config, QSendArguments, QReceiveArguments, ExchangeArguments amqp.Table) IAMQP {
	return &amqpSettings{
		SendName:             cfg.SendName,
		SendDurable:          cfg.SendDurable,
		SendDeleteWhenUnused: cfg.SendDeleteWhenUnused,
		SendExclusive:        cfg.SendQueueExclusive,
		SendNoWait:           cfg.SendQueueNoWait,
		QSendArguments:       QSendArguments,

		ReceiveName:             cfg.ReceiveName,
		ReceiveDurable:          cfg.ReceiveDurable,
		ReceiveDeleteWhenUnused: cfg.ReceiveDeleteWhenUnused,
		ReceiveExclusive:        cfg.ReceiveQueueExclusive,
		ReceiveNoWait:           cfg.ReceiveQueueNoWait,
		QReceiveArguments:       QReceiveArguments,

		ExchangeName:      cfg.Exchange,
		Type:              cfg.Type,
		ExchangeDurable:   cfg.ExchangeDurable,
		AutoDelete:        cfg.AutoDelete,
		Internal:          cfg.Internal,
		ExchangeNoWait:    cfg.ExchangeNoWait,
		ExchangeArguments: ExchangeArguments,
	}
}

func (aq *amqpSettings) GetSendQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		aq.SendName,
		aq.SendDurable,
		aq.SendDeleteWhenUnused,
		aq.SendExclusive,
		aq.SendNoWait,
		aq.QSendArguments,
	)
}

func (aq *amqpSettings) GetReceiveChan(ch *amqp.Channel) (<-chan amqp.Delivery, error) {
	err := ch.ExchangeDeclare(
		aq.ExchangeName,
		aq.Type,
		aq.ExchangeDurable,
		aq.AutoDelete,
		aq.Internal,
		aq.ExchangeNoWait,
		aq.ExchangeArguments,
	)
	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		aq.ReceiveName,
		aq.ReceiveDurable,
		aq.ReceiveDeleteWhenUnused,
		aq.ReceiveExclusive,
		aq.ReceiveNoWait,
		aq.QReceiveArguments,
	)
	if err != nil {
		return nil, err
	}

	err = ch.QueueBind(
		q.Name,          // queue name
		"",              // routing key
		aq.ExchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
}
