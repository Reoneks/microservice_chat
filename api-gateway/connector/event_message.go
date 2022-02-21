package connector

import (
	"context"
	"encoding/json"

	"github.com/Reoneks/microservice_chat/api-gateway/tools"
	"github.com/Reoneks/microservice_chat/proto"
	"github.com/streadway/amqp"
)

type EventType int64

const (
	SubscribeEventType       EventType = iota + 1 // Connect to room		Data: "1","2","3"
	UnSubscribeRoomEventType                      // Disconnect from room  Data: ""
	NewMessageEventType                           // Message				Data: "Text Message"
	EditMessageEventType
	DeleteMessageEventType
	WritingMessageEventType
	GetAllRoomsEventType
	AddedToRoomEventType
	DeletedFromRoomEventType
)

type EventMessage struct {
	Type EventType   `json:"type"`
	Data interface{} `json:"data"`
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
		data := struct {
			RoomID string `json:"room_id"`
		}{}

		err := tools.Bind(message.Data, &data)
		if err != nil {
			c.log.Error(err)
			return
		}

		roomNumber := data.RoomID
		conn.SetCurrentRoom(roomNumber)
		if c.rooms[roomNumber] == nil {
			newRoom := NewRoomConnections(roomNumber)
			c.rooms[roomNumber] = &newRoom
		}
		room := *c.rooms[roomNumber]
		room.AddConnection(conn)

		msg := proto.RabbitMessage{
			Message:   []byte(roomNumber),
			EventType: int64(SubscribeEventType),
			RoomID:    roomNumber,
		}

		b, err := json.Marshal(&msg)
		if err != nil {
			c.log.Error(err)
			return
		}

		err = c.ch.Publish(
			c.exchange,
			"",
			c.mandatory,
			c.immediate,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        b,
			},
		)
		if err != nil {
			c.log.Error(err)
			return
		}
	//! //////////////////////////////////
	case UnSubscribeRoomEventType:
		data := struct {
			RoomID string `json:"room_id"`
		}{}

		err := tools.Bind(message.Data, &data)
		if err != nil {
			c.log.Error(err)
			return
		}

		roomNumber := data.RoomID
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
			UserId: conn.GetUserID(),
		}

		b, err := json.Marshal(sendToFront)
		if err != nil {
			c.log.Error(err)
			return
		}

		msg := proto.RabbitMessage{
			Message:   b,
			EventType: int64(UnSubscribeRoomEventType),
			RoomID:    sendToFront.RoomId,
		}

		b, err = json.Marshal(&msg)
		if err != nil {
			c.log.Error(err)
			return
		}

		err = c.ch.Publish(
			c.exchange,
			"",
			c.mandatory,
			c.immediate,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        b,
			},
		)
		if err != nil {
			c.log.Error(err)
			return
		}
	//! //////////////////////////////////
	case NewMessageEventType:
		data := struct {
			Text           string  `json:"text"`
			QuoteMessageId *string `json:"quote_message_id"`
		}{}

		err := tools.Bind(message.Data, &data)
		if err != nil {
			c.log.Error(err)
			return
		}

		msg := map[string]interface{}{
			"text":       data.Text,
			"status":     1,
			"room_id":    conn.GetCurrentRoom(),
			"created_by": conn.GetUserID(),
		}

		if data.QuoteMessageId != nil {
			msg["quote_message_id"] = *data.QuoteMessageId
		}

		bytes, err := json.Marshal(&msg)
		if err != nil {
			c.log.Error(err)
			return
		}

		message := proto.RabbitMessage{
			Message:     bytes,
			MessageType: "create",
			EventType:   int64(NewMessageEventType),
			RoomID:      conn.GetCurrentRoom(),
		}

		bytes, err = json.Marshal(&message)
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
	//! //////////////////////////////////
	case EditMessageEventType:
		data := struct {
			Text      string `json:"text"`
			MessageID string `json:"message_id"`
		}{}

		err := tools.Bind(message.Data, &data)
		if err != nil {
			c.log.Error(err)
			return
		}

		msg := map[string]interface{}{
			"_id":    data.MessageID,
			"text":   data.Text,
			"status": 3,
		}

		bytes, err := json.Marshal(&msg)
		if err != nil {
			c.log.Error(err)
			return
		}

		message := proto.RabbitMessage{
			Message:     bytes,
			MessageType: "update",
			EventType:   int64(EditMessageEventType),
			RoomID:      conn.GetCurrentRoom(),
		}

		bytes, err = json.Marshal(&message)
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
	//! //////////////////////////////////
	case DeleteMessageEventType:
		data := struct {
			MessageID string `json:"message_id"`
		}{}

		err := tools.Bind(message.Data, &data)
		if err != nil {
			c.log.Error(err)
			return
		}

		message := proto.RabbitMessage{
			Message:     []byte(data.MessageID),
			MessageType: "delete",
			EventType:   int64(DeleteMessageEventType),
			RoomID:      conn.GetCurrentRoom(),
		}

		bytes, err := json.Marshal(&message)
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
	//! //////////////////////////////////
	case WritingMessageEventType:
		sendToFront := struct {
			RoomId string `json:"room_id"`
			UserId string `json:"user"`
		}{
			RoomId: conn.GetCurrentRoom(),
			UserId: conn.GetUserID(),
		}

		b, err := json.Marshal(sendToFront)
		if err != nil {
			c.log.Error(err)
			return
		}

		msg := proto.RabbitMessage{
			Message:   b,
			EventType: int64(WritingMessageEventType),
			RoomID:    sendToFront.RoomId,
		}

		b, err = json.Marshal(&msg)
		if err != nil {
			c.log.Error(err)
			return
		}

		err = c.ch.Publish(
			c.exchange,
			"",
			c.mandatory,
			c.immediate,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        b,
			},
		)
		if err != nil {
			c.log.Error(err)
			return
		}
	//! //////////////////////////////////
	case GetAllRoomsEventType:
		data := struct {
			Limit  int64 `json:"limit"`
			Offset int64 `json:"offset"`
		}{}

		err := tools.Bind(message.Data, &data)
		if err != nil {
			c.log.Error(err)
			return
		}

		if data.Limit <= 0 {
			data.Limit = c.defaltLimit
		}

		if data.Offset < 0 {
			data.Offset = c.defaltOffset
		}

		req := &proto.GetAllRoomsRequest{
			Limit:  data.Limit,
			Offset: data.Offset,
			UserID: conn.GetUserID(),
		}

		resp, err := c.room.GetAllRooms(context.Background(), req)
		if err != nil {
			c.log.Error(err)
			return
		}

		var sendToFront []map[string]interface{}
		err = json.Unmarshal(resp.Room, &sendToFront)
		if err != nil {
			c.log.Error(err)
			return
		}

		message := EventMessage{
			Type: WritingMessageEventType,
			Data: sendToFront,
		}
		conn.SendMessage(message)
	}
}
