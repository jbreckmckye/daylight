package new

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"


)

// IPInfo is the data we get from the IPInfo API
type IPInfo struct {
	IP      string `json:"ip"`
	Loc     string `json:"loc"`
	TZ      string `json:"timezone"`
}

const ipinfoUrl = "https://ipinfo.io/json?inc=ip,loc,timezone"

// FetchIPInfo makes the http call
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

// Config turns an IPInfo result into a Config, performing validation
func (inf *IPInfo) Config() (cfg Config, err error) {
	// API Loc -> config Latitude, Longitude
  if inf.Loc == "" {
		return cfg, fmt.Errorf("IPInfo did not return location data")
	}

	latlong, err := locationToLatLong(inf.Loc)
	if err != nil {
		return cfg, fmt.Errorf("IPInfo returned invalid lat/long: %w", err)
	}

	cfg.Latitude = &latlong.Lat
	cfg.Longitude = &latlong.Lng
	
	// API TZ -> config Timezone
	if inf.TZ == "" {
		return cfg, fmt.Errorf("IPInfo did not return timezone data")
	}

	tz, err := time.LoadLocation(inf.TZ)
	if err != nil {
		return cfg, fmt.Errorf("IPInfo returned invalid timezone: %w", err)
	}

	cfg.Timezone = tz

	// API IP -> config IP
	cfg.IP = &inf.IP

	return cfg, nil
}

func locationToLatLong(loc string) (LatLong, error) {
	result := LatLong{}
	parseError := fmt.Errorf("cannot parse format of location data %q", loc)

	parts := strings.Split(loc, ",")
	if len(parts) != 2 {
		return result, parseError
	}

	lat, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return result, parseError
	}

	if lat < -90 || lat > 90 {
		return result, fmt.Errorf("latitude must be between -90 and 90, was %f", lat)
	}

	lng, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return result, parseError
	}

	if lng < -180 || lng > 180 {
		return result, fmt.Errorf("longitude must be between -180 and 180, was %f", lng)
	}

	result.Lat = lat
	result.Lng = lng
	return result, nil
}