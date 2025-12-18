package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.New()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		panic(err)
	}

	r.GET("/", func(c *gin.Context) {
		remoteAddr := c.Request.RemoteAddr
		ip := formatIp(remoteAddr)
		c.String(http.StatusOK, ip)
	})

	r.Run("0.0.0.0:8080")

}

func formatIp(ip string) string {
	if bracketPos := strings.LastIndex(ip, ":"); bracketPos != -1 {
		ip = ip[:bracketPos]

	}

	ip = strings.TrimLeft(ip, "[")
	ip = strings.TrimRight(ip, "]")

	return ip
}
