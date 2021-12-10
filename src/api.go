package main

import (
	"net"

	"github.com/gin-gonic/gin"
)

func getIP(ip string, lang string) {

}

func getLocationIP(cache *CacheMU) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		ip := c.Query("ip")
		language := c.DefaultQuery("lang", "")

		if ip == "" {
			c.JSON(400, gin.H{
				"message": "Need Ip address in params",
			})
		} else {
			_, ipv4Net, err := net.ParseCIDR(ip)

			if err != nil {
				c.JSON(400, gin.H{
					"message": "Need valid IP address in format CIDR",
				})
				return
			}

			getIP(ipv4Net.String(), language)
		}
	}

	return gin.HandlerFunc(fn)
}

func putLocation(cache *CacheMU) gin.HandlerFunc {
	fn := func(c *gin.Context) {

	}

	return gin.HandlerFunc(fn)
}
