package main

import (
	"net"

	"github.com/gin-gonic/gin"
)

func getIpLocationAPI(cache *CacheMU, ip string, lang string) (IpLocationAPI, error) {
	toReturn := IpLocationAPI{
		Ip:        "",
		Locations: nil,
	}

	ipLoc, err := findIP(cache, ip)

	if err != nil {
		return toReturn, err
	}
	if ipLoc.Ip == "" {
		return toReturn, nil
	}
	toReturn.Ip = ipLoc.Ip

	loc, err := findLocation(cache, ipLoc.Uuid, lang)

	if err != nil {
		return toReturn, err
	}
	if loc == nil {
		return toReturn, nil
	}
	toReturn.Locations = loc

	addResultInCache(cache, ipLoc, loc)

	return toReturn, nil
}

func getLocationIP(cache *CacheMU) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		ip := c.Query("ip")
		language := c.Query("lang")

		if language != "EN" && language != "ES" && language != "FR" && language != "" {
			c.JSON(400, gin.H{
				"message": "Need valid language: EN, ES, FR or nothing",
			})
			return
		}

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

			result, err := getIpLocationAPI(cache, ipv4Net.String(), language)

			if err != nil {
				c.JSON(500, gin.H{
					"message": err.Error(),
				})
				return
			}

			if result.Ip == "" {
				c.JSON(404, gin.H{
					"message": "Cannot find IP",
				})
				return
			}

			if result.Locations == nil {
				c.JSON(404, gin.H{
					"message": "Cannot find Location of the ID",
				})
				return
			}

			c.JSON(200, result)
		}
	}

	return gin.HandlerFunc(fn)
}

func putLocation(cache *CacheMU) gin.HandlerFunc {
	fn := func(c *gin.Context) {

	}

	return gin.HandlerFunc(fn)
}
