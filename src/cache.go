package main

func findLocation(cache *CacheMU, uuid string, lang string) ([]LocationsLanguage, error) {
	cache.LockLocations.RLock()
	defer cache.LockLocations.RUnlock()
	var value []LocationsLanguage = nil
	var check bool = false

	if cache.Cache.Locations != nil {
		//value, check = cache.Cache.Locations[uuid]
	}

	if check == false {

	}
	return value, nil
}

func findIP(cache *CacheMU, ip string) (IpLocation, error) {
	cache.LockIP.RLock()
	defer cache.LockIP.RUnlock()
	value := IpLocation{Ip: "", Uuid: ""}
	var check bool = false

	if cache.Cache.Locations != nil {
		value, check = cache.Cache.Ip[ip]
	}

	if check == false {

	}
	return value, nil
}
