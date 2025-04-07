package new

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nathan-osman/go-sunrise"
)

type LatLong struct {
	Lat float64
	Lng float64
}

type SunTimes struct {
	// These times are only valid if !PolarDay && !PolarNight
	Rises  time.Time
	Sets   time.Time
	Length time.Duration
	// 24-hour day / night at far latitudes
	PolarNight bool
	PolarDay   bool
}

var nilTime = time.Time{}

func LocationToLatLong(loc string) (LatLong, error) {
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

// SunTimesForPlaceDate returns the sun rise / set times for the given lat/long in UTC.
// It also checks whether we are in polar day / polar night.
func SunTimesForPlaceDate(latlong LatLong, date time.Time) SunTimes {
	year, month, day := date.Date()

	rises, sets := sunrise.SunriseSunset(
		latlong.Lat, latlong.Lng,
		year, month, day,
	)

	length := sets.Sub(rises)

	polarDay := false
	polarNight := false

	// go-sunrise returns empty time.Time{} values if in polar day / night
	if rises == nilTime && sets == nilTime {
		isNorth := latlong.Lat >= 0
		isSouth := !isNorth

		isSummer := (month > time.March) && (month < time.October)
		isWinter := !isSummer

		switch {
		case isNorth && isSummer, isSouth && isWinter:
			{
				polarDay = true
				length = time.Hour * 24
			}

		case isNorth && isWinter, isSouth && isSummer:
			{
				polarNight = true
				length = time.Duration(0)
			}
		}
	}

	return SunTimes{
		Rises:      rises,
		Sets:       sets,
		Length:     length,
		PolarNight: polarNight,
		PolarDay:   polarDay,
	}
}

func SunTimesYesterday(latlong LatLong, today time.Time) SunTimes {
	yesterday := today.AddDate(0, 0, -1)
	return SunTimesForPlaceDate(latlong, yesterday)
}

func SunTimesForward(latlong LatLong, today time.Time, count int) ([]time.Time, []SunTimes) {
	projections := make([]SunTimes, count)
	dates := make([]time.Time, count)

	for i := 0; i < count; i++ {
		date := today.AddDate(0, 0, i+1)
		dates[i] = date
		projections[i] = SunTimesForPlaceDate(latlong, date)
	}
	return dates, projections
}

func (s SunTimes) ApproximateNoon() time.Time {
	return s.Rises.Add(s.Length / 2)
}
