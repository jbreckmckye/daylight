package internal

import (
	"strconv"
	"time"
)

type TodayView struct {
	Lat           string
	Lng           string
	Rise          string
	Sets          string
	Noon          string
	IP            string
	Len           string
	Diff          string
	DayStartRatio float64
	DayEndRatio   float64
	Next10Days    []DayProjection
}

type DayProjection struct {
	Day    string
	Rise   string
	Sets   string
	Length string
}

func TodayStats(today time.Time, timezone *time.Location, latlong LatLong, IP string) TodayView {
	sunTimes := SunTimesForPlaceDate(latlong, today)
	sunTimesYesterday := SunTimesYesterday(latlong, today)

	dayStartRatio, dayEndRatio := FormatDayRatio(sunTimes, timezone)

	viewmodel := TodayView{
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
		Next10Days:    projections(today, timezone, latlong, 10),
	}

	return viewmodel
}

func projections(today time.Time, timezone *time.Location, latlong LatLong, count int) []DayProjection {
	projectedDates, projectedSunTimes := SunTimesForward(latlong, today, count)
	outputs := make([]DayProjection, count)

	for i, date := range projectedDates {
		sunTimes := projectedSunTimes[i]
		outputs[i] = DayProjection{
			Day:    FormatDate(date),
			Rise:   FormatRises(sunTimes, timezone),
			Sets:   FormatSets(sunTimes, timezone),
			Length: FormatDayLength(sunTimes),
		}
	}

	return outputs
}
