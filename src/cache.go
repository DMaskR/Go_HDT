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
		}
	}()
}

func location(cache *CacheMU, uuid string, lang string) (LocationsLanguageAPI, error) {
	locate := Location{}
	value := LocationsLanguageAPI{Name: lang, Locations: locate}
	var check bool = false

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
			return value, err
		}
	}
	value.Name = lang
	value.Locations = locate
	return value, nil
}

func locationAsync(wg *sync.WaitGroup, cache *CacheMU, uuid string, lang string) (chan LocationsLanguageAPI, chan error) {
	resultChan := make(chan LocationsLanguageAPI, 1)
	errorChan := make(chan error, 1)

	wg.Add(1)

	go func(wg *sync.WaitGroup, cache *CacheMU, uuid string, lang string, resultChan chan LocationsLanguageAPI, errorChan chan error) {
		defer wg.Done()
		defer close(resultChan)
		defer close(errorChan)

		result, err := location(cache, uuid, lang)
		if err != nil {
			errorChan <- err
		} else {
			resultChan <- result
			errorChan <- nil
		}

	}(wg, cache, uuid, lang, resultChan, errorChan)
	return resultChan, errorChan
}

func findLocation(cache *CacheMU, uuid string, lang string) ([]LocationsLanguageAPI, error) {
	cache.LockLocations.RLock()
	defer cache.LockLocations.RUnlock()

	toReturn := make([]LocationsLanguageAPI, 0)

	if lang == "" {
		wg := &sync.WaitGroup{}

		location2Chan, error2Chan := locationAsync(wg, cache, uuid, "ES")

		location3Chan, error3Chan := locationAsync(wg, cache, uuid, "EN")

		location4Chan, error4Chan := locationAsync(wg, cache, uuid, "FR")

		wg.Wait()

		err2 := <-error2Chan
		err3 := <-error3Chan
		err4 := <-error4Chan

		if err2 != nil {
			return toReturn, err2
		} else if err3 != nil {
			return toReturn, err3
		} else if err4 != nil {
			return toReturn, err4
		}

		toReturn = append(toReturn, <-location2Chan)
		toReturn = append(toReturn, <-location3Chan)
		toReturn = append(toReturn, <-location4Chan)

	} else {
		value, err := location(cache, uuid, lang)
		if err != nil {
			return nil, err
		}
		toReturn = append(toReturn, value)
	}

	return toReturn, nil
}

func findIP(cache *CacheMU, ip string) (IpLocation, error) {
	cache.LockIP.RLock()
	defer cache.LockIP.RUnlock()
	value := IpLocation{Ip: "", Uuid: ""}
	var check bool = false

	if cache.Cache.Locations != nil {
		value, check = cache.Cache.Ip[ip]
	}

	if !check {
		var err error = nil
		value, err = findIpInFile(ip)
		if err != nil {
			return value, err
		}
	}
	return value, nil
}
