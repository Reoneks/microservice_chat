package config

import "github.com/go-chi/jwtauth"

func (c *Config) NewJWT() *jwtauth.JWTAuth {
	return jwtauth.New(c.Algorithm, []byte(c.Secret), nil)
}
