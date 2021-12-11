package main

func findLocationInFile(uuid string, lang string) (Location, error) {
	locate := Location{}
	return locate, nil
}

func findIpInFile(ip string) (IpLocation, error) {
	ipLoc := IpLocation{}
	return ipLoc, nil
}

/* func loadAllData(src string, dest string) (AllData, error) {

	pathFile := path.Join(dest, "IP-locations")

	os.RemoveAll(pathFile)

	allData := AllData{
		Ip:        nil,
		Locations: nil,
	}

	err := archiver.Unarchive(src, dest)

	if err != nil {
		return allData, err
	}

	wg := &sync.WaitGroup{}

	_, error1Chan := parseIpLocationAsync(wg, path.Join(pathFile, "IP-locations.csv"))

	_, error2Chan := parseLocationAsync(wg, path.Join(pathFile, "Locations-FR.csv"), "FR")

	_, error3Chan := parseLocationAsync(wg, path.Join(pathFile, "Locations-EN.csv"), "EN")

	_, error4Chan := parseLocationAsync(wg, path.Join(pathFile, "Locations-ES.csv"), "ES")

	wg.Wait()

	err1 := <-error1Chan
	err2 := <-error2Chan
	err3 := <-error3Chan
	err4 := <-error4Chan

	if err1 != nil {
		return allData, err1
	} else if err2 != nil {
		return allData, err2
	} else if err3 != nil {
		return allData, err3
	} else if err4 != nil {
		return allData, err4
	}

	allData.Ip = <-ipLocationChan
	allData.Locations = append(allData.Locations, <-location1Chan)
	allData.Locations = append(allData.Locations, <-location2Chan)
	allData.Locations = append(allData.Locations, <-location3Chan)

	return allData, nil
}

func parseIpLocationAsync(wg *sync.WaitGroup, src string) (chan []IpLocation, chan error) {
	resultChan := make(chan []IpLocation, 1)
	errorChan := make(chan error, 1)

	wg.Add(1)

	go func(wg *sync.WaitGroup, src string, resultChan chan []IpLocation, errorChan chan error) {
		defer wg.Done()
		defer close(resultChan)
		defer close(errorChan)

		result, err := parseIpLocation(src)

		if err != nil {
			errorChan <- err
		} else {
			resultChan <- result
			errorChan <- nil
		}
	}(wg, src, resultChan, errorChan)
	return resultChan, errorChan
}

func parseIpLocation(src string) ([]IpLocation, error) {
	csvFile, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	csvReader.Comma = ','

	csvLines, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var toReturn = make([]IpLocation, 0)

	for i, line := range csvLines {
		if i == 0 {
			continue
		}

		temp := IpLocation{
			Ip:   line[0],
			Uuid: line[1],
		}
		toReturn = append(toReturn, temp)
	}

	return toReturn, err
}

func parseLocationAsync(wg *sync.WaitGroup, src string, language string) (chan LocationsLanguage, chan error) {
	resultChan := make(chan LocationsLanguage, 1)
	errorChan := make(chan error, 1)

	wg.Add(1)

	go func(wg *sync.WaitGroup, src string, language string, resultChan chan LocationsLanguage, errorChan chan error) {
		defer wg.Done()
		defer close(resultChan)
		defer close(errorChan)

		_, err := parseLocation(src)

		if err != nil {
			errorChan <- err
		} else {
			resultChan <- LocationsLanguage{Name: language, Locations: result}
			errorChan <- nil
		}
	}(wg, src, language, resultChan, errorChan)
	return resultChan, errorChan
}

func parseLocation(src string) ([]Location, error) {
	csvFile, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	csvReader.Comma = ';'

	csvLines, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var toReturn = make([]Location, 0)

	for i, line := range csvLines {
		if i == 0 {
			continue
		}

		temp := Location{
			Uuid:       line[0],
			Continent:  line[1],
			Country:    line[2],
			Region:     line[3],
			Department: line[4],
			City:       line[5],
		}
		toReturn = append(toReturn, temp)
	}

	return toReturn, err
}
*/
