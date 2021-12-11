package main

func getCacheLocation(cache *CacheMU, uuid string, lang string) {

}

func getCacheIP(cache *CacheMU, ip string, lang string) {
	cache.LockIP.RLock()
	defer cache.LockIP.RUnlock()

	value, check := cache.Cache.Ip[ip]

	if check == false {

	}
	getCacheLocation(cache, value.Uuid, lang)
}
