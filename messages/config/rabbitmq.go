package config

import "github.com/streadway/amqp"

func StartRabbitMQ(url string) (*amqp.Channel, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	return conn.Channel()
}

type AmqpSettings struct {
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
}

func (s *AmqpSettings) SetQoS(ch *amqp.Channel) error {
	return ch.Qos(
		s.PrefetchCount,
		s.PrefetchSize,
		s.Global,
	)
}

func (aq *AmqpSettings) GetQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		aq.Name,
		aq.Durable,
		aq.DeleteWhenUnused,
		aq.Exclusive,
		aq.NoWait,
		aq.Arguments,
	)
}
