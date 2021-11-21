package server

import (
	"net/url"

	"github.com/Reoneks/microservice_chat/api-gateway/clients"
	"github.com/Reoneks/microservice_chat/api-gateway/server/http"
	"github.com/Reoneks/microservice_chat/proto"

	goMicroHttp "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
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
	auth                    proto.AuthService
	apiGatewaySubscribeName string
}

func NewHTTPServer(
	log *logrus.Logger,
	url *url.URL,
	authMicroservice *clients.AuthMicroservice,
	userMicroservice *clients.UserMicroservice,
	roomMicroservice *clients.RoomMicroservice,
	auth proto.AuthService,
	apiGatewaySubscribeName string,
) HTTPServer {
	return &httpServer{
		log:                     log,
		url:                     url,
		authMicroservice:        authMicroservice,
		userMicroservice:        userMicroservice,
		roomMicroservice:        roomMicroservice,
		apiGatewaySubscribeName: apiGatewaySubscribeName,
		auth:                    auth,
	}
}

func (s *httpServer) Start() error {
	srv := goMicroHttp.NewServer(
		server.Name("Api-Gateway"),
		server.Address(s.url.Host),
	)

	router := echo.New()
	router.Use(http.LoggerMiddleware())
	router.Use(http.CorsMiddleware())
	router.Use(middleware.Recover())

	router.POST("/registration", s.authMicroservice.Register)
	router.POST("/login", s.authMicroservice.Login)

	private := router.Group("/client")
	private.Use(http.Authorization(s.auth))
	{
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

	hd := srv.NewHandler(router)
	if err := srv.Handle(hd); err != nil {
		return err
	}

	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(registry.NewRegistry()),
	)
	service.Init()
	return service.Run()
}
