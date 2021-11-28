package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type service struct {
	Provider
}

func (s *service) Run() error {
	r := gin.Default()
	r.GET("/", statusHandler)
	return r.Run(":8000")
}

func statusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"gateway": http.StatusText(http.StatusOK),
	})
}
