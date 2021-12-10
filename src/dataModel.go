package main

type AllData struct {
	Ip        []IpLocation
	Locations []LocationsLanguage
}

type IpLocation struct {
	Ip   string
	Uuid string
}

type LocationsLanguage struct {
	Name      string
	Locations []Location
}

type Location struct {
	Uuid       string `json:"uuid"`
	Continent  string `json:"continent"`
	Country    string `json:"country"`
	Region     string `json:"region"`
	Department string `json:"department"`
	City       string `json:"city"`
}
