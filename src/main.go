package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func AllDataToCacheMu(data AllData) CacheMU {
	return CacheMU{LockIP: sync.RWMutex{}, LockLocations: sync.RWMutex{}, Cache: data}
}

func main() {

	allData := AllData{
		Ip:        make(map[string]IpLocation),
		Locations: make(map[string]LocationsLanguage),
	}

	cache := AllDataToCacheMu(allData)

	gin.DisableConsoleColor()

	f, _ := os.Create("trace.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		var toReturn string = ""

		toReturn = fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)

		if param.StatusCode == 200 {
			toReturn = toReturn + " Ip: " + param.Keys["IP"].(string) + " Loc:"

			for key, value := range param.Keys["Loc"].(map[string]string) {
				toReturn = toReturn + " " + key + "=" + value
			}
		}
		return toReturn + "\n"
	}))
	router.Use(gin.Recovery())

	router.GET("/location", getLocationIP(&cache))
	router.PUT("/locations", putLocation(&cache))

	router.Run("localhost:8080")
}
