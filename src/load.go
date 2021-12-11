package main

import (
	"encoding/csv"
	"io"

	"github.com/mholt/archiver"
)

func findLocationInFile(uuid string, lang string) (Location, error) {
	locate := Location{}
	err := archiver.Walk("IP-locations.rar", func(f archiver.File) error {
		var err error = nil
		if f.Name() == "IP-locations/Locations-"+lang+".csv" {
			csvReader := csv.NewReader(f.ReadCloser)
			csvReader.Comma = ';'

			for err == nil {
				csvLines, err := csvReader.Read()
				if err == io.EOF {
					return nil
				}
				if err != nil {
					return err
				}
				if uuid == csvLines[0] {
					locate.Uuid = csvLines[0]
					locate.Continent = csvLines[1]
					locate.Country = csvLines[2]
					locate.Region = csvLines[3]
					locate.Department = csvLines[4]
					locate.City = csvLines[5]
					return nil
				}
			}
		}
		return err
	})
	if err != nil {
		return locate, err
	}
	return locate, nil
}

func findIpInFile(ip string) (IpLocation, error) {
	ipLoc := IpLocation{}
	err := archiver.Walk("IP-locations.rar", func(f archiver.File) error {
		var err error = nil
		if f.Name() == "IP-locations/IP-locations.csv" {
			csvReader := csv.NewReader(f.ReadCloser)
			csvReader.Comma = ','

			for err == nil {
				csvLines, err := csvReader.Read()
				if err == io.EOF {
					return nil
				}
				if err != nil {
					return err
				}
				if ip == csvLines[0] {
					ipLoc.Ip = csvLines[0]
					ipLoc.Uuid = csvLines[1]
					return nil
				}
			}
		}
		return err
	})
	if err != nil {
		return ipLoc, err
	}
	return ipLoc, nil
}
