package client

import "github.com/gin-gonic/gin"

func NewClient() *client {
	return new(client)
}

type Provider interface {
	GetStatus(*gin.Context)
}
