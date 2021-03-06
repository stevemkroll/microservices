package service

import (
	"fmt"
	"net/http"
	notification "notification/client"

	"github.com/gin-gonic/gin"
)

type service struct {
	Provider
}

func (s *service) Run() error {
	r := gin.Default()
	r.GET("/", statusHandler)

	notificationCLient := notification.NewClient()
	r.GET("/notification", notificationCLient.GetStatus)

	return r.Run(":8000")
}

func statusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"gateway": fmt.Sprintf("%d %s",
			http.StatusOK,
			http.StatusText(http.StatusOK),
		),
	})
}
