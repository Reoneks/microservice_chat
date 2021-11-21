package http

import (
	"context"
	"net/http"
	"strings"

	"github.com/Reoneks/microservice_chat/proto"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CorsMiddleware() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})
}

func LoggerMiddleware() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${path} | ${method} | ${status} | ${latency_human}\n",
	})
}

const BearerToken = "Bearer "

func getToken(r *http.Request, findTokenFns ...func(r *http.Request) string) string {
	tokenStr := ""
	for _, fn := range findTokenFns {
		tokenStr = fn(r)
		if tokenStr != "" {
			break
		}
	}
	return tokenStr
}

func tokenFromQuery(r *http.Request) string {
	return r.URL.Query().Get("jwt")
}

func tokenFromHeader(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	if strings.Contains(bearer, BearerToken) {
		bearer = bearer[len(BearerToken):]
	}
	return bearer
}

func Authorization(auth proto.AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			var data proto.Token
			data.Token = getToken(ctx.Request(), tokenFromQuery, tokenFromHeader)
			if data.Token == "" {
				return ctx.NoContent(http.StatusUnauthorized)
			}

			authResp, err := auth.AuthHandler(context.Background(), &data)
			if err != nil {
				return ctx.JSON(http.StatusInternalServerError, err)
			}

			ctx.Set("user_id", authResp.UserID)
			return next(ctx)
		}
	}
}
