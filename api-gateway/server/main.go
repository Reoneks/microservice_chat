package server

import (
	"net/url"

	"github.com/Reoneks/microservice_chat/api-gateway/clients"
	"github.com/Reoneks/microservice_chat/api-gateway/connector"
	"github.com/Reoneks/microservice_chat/api-gateway/server/http"
	"github.com/Reoneks/microservice_chat/proto"
	"github.com/gorilla/websocket"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

type HTTPServer interface {
	Start() error
}

type httpServer struct {
	log                     *logrus.Logger
	url                     *url.URL
	authMicroservice        *clients.AuthMicroservice
	userMicroservice        *clients.UserMicroservice
	roomMicroservice        *clients.RoomMicroservice
	messagesMicroservice    *clients.MessagesMicroservice
	auth                    proto.AuthService
	apiGatewaySubscribeName string
	connect                 connector.Connector
	upgrader                *websocket.Upgrader
}

func NewHTTPServer(
	log *logrus.Logger,
	url *url.URL,
	authMicroservice *clients.AuthMicroservice,
	userMicroservice *clients.UserMicroservice,
	roomMicroservice *clients.RoomMicroservice,
	messagesMicroservice *clients.MessagesMicroservice,
	auth proto.AuthService,
	apiGatewaySubscribeName string,
	connect connector.Connector,
	upgrader *websocket.Upgrader,
) HTTPServer {
	return &httpServer{
		log:                     log,
		url:                     url,
		authMicroservice:        authMicroservice,
		userMicroservice:        userMicroservice,
		roomMicroservice:        roomMicroservice,
		messagesMicroservice:    messagesMicroservice,
		apiGatewaySubscribeName: apiGatewaySubscribeName,
		auth:                    auth,
		connect:                 connect,
		upgrader:                upgrader,
	}
}

func (s *httpServer) Start() error {
	router := echo.New()
	router.Use(http.LoggerMiddleware())
	router.Use(http.CorsMiddleware())
	router.Use(middleware.Recover())

	router.POST("/registration", s.authMicroservice.Register)
	router.POST("/login", s.authMicroservice.Login)

	private := router.Group("/client")
	private.Use(http.Authorization(s.auth))
	{
		private.GET("/ws", clients.WSHandler(s.connect, s.upgrader))
		private.GET("/messages", s.messagesMicroservice.GetMessagesByRoom)

		private.GET("/users", s.userMicroservice.GetUsers)
		private.PUT("/user", s.userMicroservice.UpdateUser)
		private.DELETE("/user", s.userMicroservice.DeleteUser)
		private.GET("/user/:id", s.userMicroservice.GetUserByID)

		private.GET("/rooms", s.roomMicroservice.GetRooms)
		private.POST("/room", s.roomMicroservice.CreateRoom)
		private.PUT("/room", s.roomMicroservice.UpdateRoom)
		private.DELETE("/room", s.roomMicroservice.DeleteRoom)
		private.GET("/add_users", s.roomMicroservice.AddUsers)
		private.GET("/room/:id", s.roomMicroservice.GetRoom)
	}

	return router.Start(s.url.String())
}
