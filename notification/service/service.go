package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type service struct {
	Provider
}

func (s *service) Run() error {
	r := gin.Default()
	r.GET("/", statusHandler)
	return r.Run(":8001")
}

func statusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"notification": fmt.Sprintf("%d %s",
			http.StatusOK,
			http.StatusText(http.StatusOK),
		),
	})
}
