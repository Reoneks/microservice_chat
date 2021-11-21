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
	//& Queue
	Name             string     //^ messages
	Durable          bool       //^ false
	DeleteWhenUnused bool       //^ false
	Exclusive        bool       //^ true
	NoWait           bool       //^ false
	Arguments        amqp.Table //^ nil

	//& QoS
	PrefetchCount int  //^ 1
	PrefetchSize  int  //^ 0
	Global        bool //^ false

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
	SetQoS(ch *amqp.Channel) error
	GetQueue(ch *amqp.Channel) (amqp.Queue, error)
	SetExchange(ch *amqp.Channel) error
}

func GetAMQP(cfg *Config, Arguments, ExchangeArguments amqp.Table) IAMQP {
	return &amqpSettings{
		Name:             cfg.Name,
		Durable:          cfg.Durable,
		DeleteWhenUnused: cfg.DeleteWhenUnused,
		Exclusive:        cfg.QueueExclusive,
		NoWait:           cfg.QueueNoWait,
		Arguments:        Arguments,

		PrefetchCount: cfg.PrefetchCount,
		PrefetchSize:  cfg.PrefetchSize,
		Global:        cfg.Global,

		ExchangeName:      cfg.Exchange,
		Type:              cfg.Type,
		ExchangeDurable:   cfg.ExchangeDurable,
		AutoDelete:        cfg.AutoDelete,
		Internal:          cfg.Internal,
		ExchangeNoWait:    cfg.ExchangeNoWait,
		ExchangeArguments: ExchangeArguments,
	}
}

func (s *amqpSettings) SetExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		s.ExchangeName,      // name
		s.Type,              // type
		s.ExchangeDurable,   // durable
		s.AutoDelete,        // auto-deleted
		s.Internal,          // internal
		s.ExchangeNoWait,    // no-wait
		s.ExchangeArguments, // arguments
	)
}

func (s *amqpSettings) SetQoS(ch *amqp.Channel) error {
	return ch.Qos(
		s.PrefetchCount,
		s.PrefetchSize,
		s.Global,
	)
}

func (aq *amqpSettings) GetQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		aq.Name,
		aq.Durable,
		aq.DeleteWhenUnused,
		aq.Exclusive,
		aq.NoWait,
		aq.Arguments,
	)
}
