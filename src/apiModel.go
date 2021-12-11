package main

type IpLocationAPI struct {
	Ip        string                 `json:"id"`
	Locations []LocationsLanguageAPI `json:"infos"`
}

type LocationsLanguageAPI struct {
	Name      string   `json:"language"`
	Locations Location `json:"location"`
}
