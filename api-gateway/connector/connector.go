package connector

import (
	"encoding/json"
	"sync"

	"github.com/Reoneks/microservice_chat/api-gateway/config"
	"github.com/Reoneks/microservice_chat/proto"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type Connector interface {
	AddConnection(conn Connection)
	SendMessageByRoom(roomID string, data interface{})
}

type WSConnectorImpl struct {
	sync.RWMutex

	log   *logrus.Logger
	rooms map[string]*RoomConnections
	msgs  <-chan amqp.Delivery

	room         proto.RoomsService
	defaltLimit  int64
	defaltOffset int64

	ch *amqp.Channel

	//& publish
	qSendName string
	exchange  string
	mandatory bool //^ false
	immediate bool //^ false
}

func (c *WSConnectorImpl) AddConnection(conn Connection) {
	c.connect(conn)
}

func (c *WSConnectorImpl) SendMessageByRoom(roomID string, data interface{}) {
	if c.rooms[roomID] == nil {
		c.log.Error("connector: room is nil")
		return
	}
	room := *c.rooms[roomID]
	room.SendMessage(data)
}

func (c *WSConnectorImpl) connect(conn Connection) {
	err := conn.Connect()
	if err != nil {
		c.log.Error(err)
	}
	c.listen(conn)
}

func (c *WSConnectorImpl) disconnect(conn Connection) {
	if c.rooms[conn.GetCurrentRoom()] == nil {
		c.log.Error("connector: room is nil")
		return
	}
	room := *c.rooms[conn.GetCurrentRoom()]
	room.RemoveConnection(conn)
}

func (c *WSConnectorImpl) listen(conn Connection) {
loop:
	for {
		select {
		case msg := <-conn.GetMessageChan():
			c.onEventMessage(conn, msg)
		case <-conn.GetDisconnectChan():
			c.disconnect(conn)
			break loop
		case msg := <-c.msgs:
			var message proto.RabbitMessage

			err := json.Unmarshal(msg.Body, &message)
			if err != nil {
				c.log.Error(err)
				continue
			}

			var messageE EventMessage
			if message.EventType != int64(DeleteMessageEventType) {
				msgData := make(map[string]interface{})
				err = json.Unmarshal(message.Message, &msgData)
				if err != nil {
					c.log.Error(err)
					continue
				}

				messageE = EventMessage{
					Type: EventType(message.EventType),
					Data: msgData,
				}
			} else {
				messageE = EventMessage{
					Type: EventType(message.EventType),
					Data: map[string]interface{}{"id": string(message.Message)},
				}
			}

			c.SendMessageByRoom(message.RoomID, messageE)
		}
	}
}

func NewWSConnector(
	log *logrus.Logger,
	ch *amqp.Channel,
	cfg *config.Config,
	msgs <-chan amqp.Delivery,
	room proto.RoomsService,
) Connector {
	return &WSConnectorImpl{
		log:          log,
		rooms:        map[string]*RoomConnections{},
		room:         room,
		defaltLimit:  cfg.DefaltLimit,
		defaltOffset: cfg.DefaltOffset,
		msgs:         msgs,
		ch:           ch,
		qSendName:    cfg.SendName,
		exchange:     cfg.Exchange,
		mandatory:    cfg.Mandatory,
		immediate:    cfg.Immediate,
	}
}
