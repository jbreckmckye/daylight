package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IPInfo struct {
	IP      string `json:"ip"`
	// City    string `json:"city"`
	// Country string `json:"country"`
	Loc     string `json:"loc"`
	TZ      string `json:"timezone"`
}

const ipinfoUrl = "https://ipinfo.io/json?inc=ip,loc,timezone"

func FetchIPInfo() (IPInfo, error) {
	res, err := http.Get(ipinfoUrl)
	if err != nil {
		return IPInfo{}, fmt.Errorf("error fetching IP info from %s, error was %q", ipinfoUrl, err.Error())
	}

	defer func() {
		res.Body.Close()
	}()

	decoder := json.NewDecoder(res.Body)
	result := IPInfo{}

	err = decoder.Decode(&result)
	return result, err
}

// var cache *IPInfo

// func CachedIPInfo() (IPInfo, error) {
//   if cache != nil {
// 		return *cache, nil
// 	}

// 	ipinfo, err := FetchIPInfo()
// 	if err != nil {
// 		return IPInfo{}, err
// 	}

// 	cache = &ipinfo
// 	return ipinfo, nil
// }
