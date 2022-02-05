package http

import (
	"net/http"

	"github.com/Reoneks/microservice_chat/api-gateway/connector"
	"github.com/labstack/echo/v4"

	"github.com/gorilla/websocket"
)

func WSHandler(connect connector.Connector, upgrader *websocket.Upgrader) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		conn, err := upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
		if err != nil {
			return ctx.NoContent(http.StatusInternalServerError)
		}

		userCtx := ctx.Get("user_id")
		connection := connector.NewWSConnection(ctx.Request(), conn, userCtx.(string))
		connect.AddConnection(connection)
		return ctx.NoContent(http.StatusOK)
	}
}
