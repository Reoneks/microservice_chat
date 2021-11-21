package connector

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type Connector interface {
	AddConnection(conn Connection)
	SendMessageByRoom(roomID string, data interface{})
}

type WSConnectorImpl struct {
	sync.RWMutex

	log   *logrus.Entry
	rooms map[string]*RoomConnections
}

func (c *WSConnectorImpl) AddConnection(conn Connection) {
	c.connect(conn)
}

func (c *WSConnectorImpl) SendMessageByRoom(roomID string, data interface{}) {
	if c.rooms[roomID] == nil {
		c.log.Error("connector 32: room is nil")
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
	go c.listen(conn)
}

func (c *WSConnectorImpl) disconnect(conn Connection) {
	if c.rooms[conn.GetCurrentRoom()] == nil {
		c.log.Error("connector 32: room is nil")
		return
	}
	room := *c.rooms[conn.GetCurrentRoom()]
	room.RemoveConnection(conn)
}

func (c *WSConnectorImpl) listen(conn Connection) {
	for {
		select {
		case msg := <-conn.GetMessageChan():
			c.onEventMessage(conn, msg)
		case <-conn.GetDisconnectChan():
			c.disconnect(conn)
		}
	}
}

func NewWSConnector(log *logrus.Entry) Connector {
	return &WSConnectorImpl{
		log:   log,
		rooms: map[string]*RoomConnections{},
	}
}
