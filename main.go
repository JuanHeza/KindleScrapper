package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gin-gonic/gin"
)

var ()

func main() {
	fmt.Printf("%v:8080\n", GetOutboundIP())
	router := gin.Default()
    router.LoadHTMLGlob("templates/*")
	router.GET("/", homePage)
	router.GET("/albums", getContents)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/content", saveContent)

	router.Run(":8080")
}

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
