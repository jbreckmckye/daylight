package internal

import (
	"strconv"
)

type SummaryView struct {
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

func Summary(query DaylightQuery) SummaryView {
	location := LatLong{
		Lat: query.Lat,
		Lng: query.Long,
	}

	timezone := &query.TZ

	today := SunTimesForPlaceDate(
		location,
		query.Date,
	)

	dayStartRatio, dayEndRatio := FormatDayRatio(today, timezone)

	yesterday := SunTimesForPlaceDate(
		location,
		query.Date.AddDate(0, 0, -1),
	)

	projectedDates := make([]DayProjection, 10)
	for i := range 10 {
		date := query.Date.AddDate(0, 0, 1 + i)
		sunTimes := SunTimesForPlaceDate(
			location,
			date,
		)

		projectedDates[i] = DayProjection{
			Day:    FormatDate(date),
			Rise:   FormatRises(sunTimes, timezone),
			Sets:   FormatSets(sunTimes, timezone),
			Length: FormatDayLength(sunTimes),
		}
	}

	return SummaryView{
		Lat:           strconv.FormatFloat(query.Lat, 'g', 4, 64),
		Lng:           strconv.FormatFloat(query.Long, 'g', 4, 64),
		Rise:          FormatRises(today, timezone),
		Sets:          FormatSets(today, timezone),
		Noon:          FormatNoon(today, timezone),
		IP:            query.IP,
		Len:           FormatDayLength(today),
		Diff:          FormatLengthDiff(today, yesterday),
		DayStartRatio: dayStartRatio,
		DayEndRatio:   dayEndRatio,
		Next10Days:    projectedDates,
	}
}
