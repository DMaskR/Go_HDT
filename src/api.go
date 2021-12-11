package main

import (
	"net"

	"github.com/gin-gonic/gin"
)

func getLocationIP(cache *CacheMU) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		ip := c.Query("ip")
		language := c.Query("lang")

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

			getCacheIP(cache, ipv4Net.String(), language)
		}
	}

	return gin.HandlerFunc(fn)
}

func putLocation(cache *CacheMU) gin.HandlerFunc {
	fn := func(c *gin.Context) {

	}

	return gin.HandlerFunc(fn)
}
