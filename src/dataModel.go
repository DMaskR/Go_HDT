package main

import (
	"sync"
)

type CacheMU struct {
	LockIP        sync.RWMutex
	LockLocations sync.RWMutex
	Cache         AllData
}

type AllData struct {
	Ip        map[string]IpLocation
	Locations map[string]LocationsLanguage
}

type IpLocation struct {
	Ip   string
	Uuid string
}

type LocationsLanguage struct {
	Name      string
	Locations map[string]Location
}

type Location struct {
	Uuid       string `json:"uuid"`
	Continent  string `json:"continent"`
	Country    string `json:"country"`
	Region     string `json:"region"`
	Department string `json:"department"`
	City       string `json:"city"`
}
