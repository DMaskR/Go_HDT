package main

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

func findLocation(cache *CacheMU, uuid string, lang string) ([]LocationsLanguageAPI, error) {
	cache.LockLocations.RLock()
	defer cache.LockLocations.RUnlock()

	toReturn := make([]LocationsLanguageAPI, 0)

	if lang == "" {
		value, err := location(cache, uuid, "EN")
		if err != nil {
			return nil, err
		}
		toReturn = append(toReturn, value)

		value, err = location(cache, uuid, "ES")
		if err != nil {
			return nil, err
		}
		toReturn = append(toReturn, value)

		value, err = location(cache, uuid, "FR")
		if err != nil {
			return nil, err
		}
		toReturn = append(toReturn, value)

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
