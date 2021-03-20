package main

import (
	"fmt"
	"os"
	"net"
	"errors"

	"github.com/bgguna/me/contact"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	log "github.com/sirupsen/logrus"
)

const (
	port = "3000"
)

func init() {
	gin.SetMode(gin.ReleaseMode)

	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {
	log.Info("Server started for bgguna.")

	router := gin.Default()
	router.Use(cors.Default())

	// This is for handling CONTACT submissions.
	router.POST("/send", contact.HandleNewMsg())

	hostname, err := getLocalIP()
	if err != nil {
		log.Error("Failed to get IP address. Using localhost instead 😢...")
		hostname = "127.0.0.1"
	}

	hostname = "127.0.0.1"

	log.Infof("Listening on http://%s:%s...", hostname, port)
	router.Run(fmt.Sprintf("%s:%s", hostname, port))
}

// getLocalIP returns the local IP of the host
func getLocalIP() (string, error) {
    addrs, err := net.InterfaceAddrs()
    if err != nil {
        return "", err
	}
	
    for _, address := range addrs {
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                return ipnet.IP.String(), nil
            }
        }
	}
	
    return "", errors.New("ip address not found")
}
