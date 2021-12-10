package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

func AllDataToCacheMu(data AllData) CacheMU {
	return CacheMU{Lock: sync.RWMutex{}, Cache: data}
}

func main() {

	data, err := loadAllData("IP-locations.rar", "./")

	if err != nil {
		fmt.Println(err)
		return
	}

	cache := AllDataToCacheMu(data)

	gin.DisableConsoleColor()

	f, _ := os.Create("trace.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.GET("/location", getLocationIP(&cache))
	router.PUT("/locations", putLocation(&cache))

	router.Run("localhost:8080")
}
