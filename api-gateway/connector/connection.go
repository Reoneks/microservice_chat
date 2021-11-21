package connector

import (
	"net/http"

	"github.com/Reoneks/microservice_chat/api-gateway/model"
	"github.com/gorilla/websocket"
)

type Connection interface {
	GetMessageChan() chan []byte
	GetDisconnectChan() chan struct{}
	Connect() error
	Disconnect() error
	SendMessage(data interface{}) error
	ReadMessage() (messageType int, p []byte, err error)
	GetUser() *model.User
	GetCurrentRoom() string
	SetCurrentRoom(roomID string)
}

type WSConnection struct {
	r             *http.Request
	conn          *websocket.Conn
	user          *model.User
	currentRoomID string

	isConnected    bool
	messageChan    chan []byte
	disconnectChan chan struct{} // on disconnect
	closeConnChan  chan struct{} // on close
}

func (c *WSConnection) GetMessageChan() chan []byte {
	return c.messageChan
}

func (c *WSConnection) GetDisconnectChan() chan struct{} {
	return c.disconnectChan
}

func (c *WSConnection) Connect() error {
	if c.isConnected {
		return nil
	}
	c.isConnected = true
	go c.connect()
	return nil
}

func (c *WSConnection) connect() {
	for {
		select {
		case <-c.closeConnChan:
			c.isConnected = false
			return
		default:
			_, msgData, err := c.conn.ReadMessage()
			if err != nil {
				c.isConnected = false
				c.disconnectChan <- struct{}{}
				return
			}
			c.messageChan <- msgData
		}
	}

}

func (c *WSConnection) Disconnect() error {
	c.closeConnChan <- struct{}{}
	return c.conn.Close()
}

func (c *WSConnection) SendMessage(data interface{}) error {
	return c.conn.WriteJSON(data)
}

func (c *WSConnection) ReadMessage() (messageType int, p []byte, err error) {
	return c.conn.ReadMessage()
}

func (c *WSConnection) GetUser() *model.User {
	return c.user
}

func (c *WSConnection) GetCurrentRoom() string {
	return c.currentRoomID
}

func (c *WSConnection) SetCurrentRoom(roomID string) {
	c.currentRoomID = roomID
}

func NewWSConnection(r *http.Request, conn *websocket.Conn, user *model.User) Connection {
	return &WSConnection{
		r:              r,
		conn:           conn,
		user:           user,
		currentRoomID:  "",
		isConnected:    false,
		messageChan:    make(chan []byte),
		disconnectChan: make(chan struct{}),
		closeConnChan:  make(chan struct{}),
	}
}
