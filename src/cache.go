package main

import (
	"sync"
)

func addResultInCache(cache *CacheMU, ip IpLocation, loc []LocationsLanguageAPI) {
	go func() {
		cache.LockIP.Lock()
		defer cache.LockIP.Unlock()

		cache.Cache.Ip[ip.Ip] = ip
	}()

	go func() {
		cache.LockLocations.Lock()
		defer cache.LockLocations.Unlock()

		for _, element := range loc {
			lang := cache.Cache.Locations[element.Name]
			lang.Name = element.Name

			if lang.Locations == nil {
				lang.Locations = make(map[string]Location)
			}
			lang.Locations[element.Locations.Uuid] = element.Locations
			cache.Cache.Locations[element.Name] = lang
		}
	}()
}

func location(cache *CacheMU, uuid string, lang string) (LocationsLanguageAPI, string, error) {
	locate := Location{}
	value := LocationsLanguageAPI{Name: lang, Locations: locate}
	var check bool = false
	var toFind string = ""

	if cache.Cache.Locations != nil {
		valueLang, checkLang := cache.Cache.Locations[lang]
		if checkLang {
			locate, check = valueLang.Locations[uuid]
		}
	}

	if !check {
		var err error = nil
		locate, err = findLocationInFile(uuid, lang)
		if err != nil {
			return value, "", err
		}
		toFind = "file"
	} else {
		toFind = "cache"
	}
	value.Name = lang
	value.Locations = locate
	return value, toFind, nil
}

func locationAsync(wg *sync.WaitGroup, cache *CacheMU, uuid string, lang string) (chan LocationsLanguageAPI, chan string, chan error) {
	resultChan := make(chan LocationsLanguageAPI, 1)
	errorChan := make(chan error, 1)
	toFind := make(chan string, 1)

	wg.Add(1)

	go func(wg *sync.WaitGroup, cache *CacheMU, uuid string, lang string, resultChan chan LocationsLanguageAPI, toFindChan chan string, errorChan chan error) {
		defer wg.Done()
		defer close(resultChan)
		defer close(errorChan)
		defer close(toFindChan)

		result, toFind, err := location(cache, uuid, lang)
		if err != nil {
			errorChan <- err
		} else {
			toFindChan <- toFind
			resultChan <- result
			errorChan <- nil
		}

	}(wg, cache, uuid, lang, resultChan, toFind, errorChan)
	return resultChan, toFind, errorChan
}

func findLocation(cache *CacheMU, uuid string, lang string) ([]LocationsLanguageAPI, map[string]string, error) {
	cache.LockLocations.RLock()
	defer cache.LockLocations.RUnlock()

	toReturn := make([]LocationsLanguageAPI, 0)

	var toFindLoc map[string]string = make(map[string]string)

	if lang == "" {
		wg := &sync.WaitGroup{}

		location2Chan, toFind2Chan, error2Chan := locationAsync(wg, cache, uuid, "ES")

		location3Chan, toFind3Chan, error3Chan := locationAsync(wg, cache, uuid, "EN")

		location4Chan, toFind4Chan, error4Chan := locationAsync(wg, cache, uuid, "FR")

		wg.Wait()

		err2 := <-error2Chan
		err3 := <-error3Chan
		err4 := <-error4Chan

		if err2 != nil {
			return toReturn, nil, err2
		} else if err3 != nil {
			return toReturn, nil, err3
		} else if err4 != nil {
			return toReturn, nil, err4
		}

		toFindLoc["ES"] = <-toFind2Chan
		toFindLoc["EN"] = <-toFind3Chan
		toFindLoc["FR"] = <-toFind4Chan

		toReturn = append(toReturn, <-location2Chan)
		toReturn = append(toReturn, <-location3Chan)
		toReturn = append(toReturn, <-location4Chan)

	} else {
		value, toFind, err := location(cache, uuid, lang)
		if err != nil {
			return nil, nil, err
		}
		toFindLoc[lang] = toFind
		toReturn = append(toReturn, value)
	}

	return toReturn, toFindLoc, nil
}

func findIP(cache *CacheMU, ip string) (IpLocation, string, error) {
	cache.LockIP.RLock()
	defer cache.LockIP.RUnlock()
	value := IpLocation{Ip: "", Uuid: ""}
	var check bool = false
	var toFind string = ""

	if cache.Cache.Locations != nil {
		value, check = cache.Cache.Ip[ip]
	}

	if !check {
		var err error = nil
		value, err = findIpInFile(ip)
		if err != nil {
			return value, "", err
		}
		toFind = "file"
	} else {
		toFind = "cache"
	}
	return value, toFind, nil
}
