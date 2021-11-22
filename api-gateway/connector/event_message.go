package connector

import (
	"encoding/json"

	"github.com/Reoneks/microservice_chat/api-gateway/model"
	"github.com/streadway/amqp"
)

type EventType int64

const (
	SubscribeEventType       EventType = 1 // Connect to room		Data: "1","2","3"
	UnSubscribeRoomEventType EventType = 2 // Disconnect from room  Data: ""
	NewMessageEventType      EventType = 3 // Message				Data: "Text Message"
)

type EventMessage struct {
	Type EventType `json:"type"`
	Data string    `json:"data"`
}

func (c *WSConnectorImpl) onEventMessage(conn Connection, msg []byte) {
	var message EventMessage

	err := json.Unmarshal(msg, &message)
	if err != nil {
		c.log.Error(err)
		return
	}

	switch message.Type {
	//! //////////////////////////////////
	case SubscribeEventType:
		roomNumber := message.Data
		conn.SetCurrentRoom(roomNumber)
		if c.rooms[roomNumber] == nil {
			newRoom := NewRoomConnections(roomNumber)
			c.rooms[roomNumber] = &newRoom
		}
		room := *c.rooms[roomNumber]
		room.AddConnection(conn)

		sendToFront := struct {
			RoomId string
			User   *model.User
		}{
			RoomId: roomNumber,
			User:   conn.GetUser(),
		}
		bytes, err := json.Marshal(sendToFront)
		if err != nil {
			c.log.Error(err)
			return
		}
		message := EventMessage{
			Type: 1,
			Data: string(bytes),
		}
		room.SendMessage(message)
	//! //////////////////////////////////
	case UnSubscribeRoomEventType:
		roomNumber := message.Data
		pRoom := c.rooms[roomNumber]
		if pRoom == nil {
			return
		}
		room := *pRoom
		room.RemoveConnection(conn)
		sendToFront := struct {
			RoomId string
			UserId string
		}{
			RoomId: roomNumber,
			UserId: conn.GetUser().ID,
		}
		bytes, err := json.Marshal(sendToFront)
		if err != nil {
			c.log.Error(err)
			return
		}
		message := EventMessage{
			Type: 2,
			Data: string(bytes),
		}
		room.SendMessage(message)
	//! //////////////////////////////////
	case NewMessageEventType:
		message := model.Message{
			Text:        message.Data,
			Status:      1,
			RoomID:      conn.GetCurrentRoom(),
			CreatedBy:   conn.GetUser().ID,
			MessageType: "create",
		}

		bytes, err := json.Marshal(message)
		if err != nil {
			c.log.Error(err)
			return
		}

		err = c.ch.Publish(
			"",
			c.qSendName,
			c.mandatory,
			c.immediate,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(bytes),
			},
		)
		if err != nil {
			c.log.Error(err)
			return
		}
	}
}
