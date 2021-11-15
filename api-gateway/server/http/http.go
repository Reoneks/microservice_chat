package http

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	ServiceName = "Api Gateway"
)

type Response struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func VersionHandler(ctx echo.Context) error {
	version, _ := ioutil.ReadFile("VERSION")
	return ctx.JSON(http.StatusOK, Response{Name: ServiceName, Version: string(version)})
}
