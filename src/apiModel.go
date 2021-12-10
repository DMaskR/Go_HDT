package main

type IpLocationAPI struct {
	Ip        string              `json:"id"`
	Locations []LocationsLanguage `json:"infos"`
}

type LocationsLanguageAPI struct {
	Name      string   `json:"lang"`
	Locations Location `json:"location"`
}
