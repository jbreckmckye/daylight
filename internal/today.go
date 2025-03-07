package internal

import (
	"strconv"
	"time"
)

type TodayViewModel struct {
	Lat           string
	Lng           string
	Rise          string
	Sets          string
	Noon          string
	IP            string
	Len           string
	Diff          string
	DayStartRatio int
	DayEndRatio   int
}

func TodayStats(today time.Time, timezone *time.Location, latlong LatLong, IP string) TodayViewModel {
	sunTimes := SunTimesForPlaceDate(latlong, today)
	sunTimesYesterday := SunTimesYesterday(latlong, today)

	dayStartRatio, dayEndRatio := FormatDayRatio(sunTimes, timezone, 3)

	viewmodel := TodayViewModel{
		Lat:           strconv.FormatFloat(latlong.Lat, 'g', 4, 64),
		Lng:           strconv.FormatFloat(latlong.Lng, 'g', 4, 64),
		Rise:          FormatRises(sunTimes, timezone),
		Sets:          FormatSets(sunTimes, timezone),
		Noon:          FormatNoon(sunTimes, timezone),
		IP:            IP,
		Len:           FormatDayLength(sunTimes),
		Diff:          FormatLengthDiff(sunTimes, sunTimesYesterday),
		DayStartRatio: dayStartRatio,
		DayEndRatio:   dayEndRatio,
	}

	return viewmodel
}
