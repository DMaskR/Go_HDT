package main

import (
	"io"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

func AllDataToCacheMu(data AllData) CacheMU {
	return CacheMU{LockIP: sync.RWMutex{}, LockLocations: sync.RWMutex{}, Cache: data}
}

func main() {

	allData := AllData{
		Ip:        nil,
		Locations: nil,
	}

	cache := AllDataToCacheMu(allData)

	gin.DisableConsoleColor()

	f, _ := os.Create("trace.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.GET("/location", getLocationIP(&cache))
	router.PUT("/locations", putLocation(&cache))

	router.Run("localhost:8080")
}
