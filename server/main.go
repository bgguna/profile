package main

import (
	"fmt"
	"os"

	"github.com/bgguna/me/contact"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

const (
	hostname = "127.0.0.1"
	port     = "3000"
)

func init() {
	gin.SetMode(gin.ReleaseMode)

	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {
	log.Info("Server started...")

	router := gin.Default()
	router.Use(cors.Default())

	// TODO: remove; this is used only for testing.
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/send", contact.HandleNewMsg())

	log.Infof("Listening on http://%s:%s...", hostname, port)
	router.Run(fmt.Sprintf("%s:%s", hostname, port))
}
