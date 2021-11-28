package client

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type client struct {
	Provider
}

func (cl *client) GetStatus(c *gin.Context) {
	endpoint := "http://notification:8001"
	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"notification": fmt.Sprintf("%d %s",
				http.StatusServiceUnavailable,
				http.StatusText(http.StatusServiceUnavailable),
			),
		})
		return
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("\n\n%+v\n\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"notification": fmt.Sprintf("%d %s",
				http.StatusBadRequest,
				http.StatusText(http.StatusBadRequest),
			),
		})
		return
	}
	c.JSON(res.StatusCode, gin.H{
		"notification": res.Status,
	})
}
