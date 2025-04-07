package new

import (
	"fmt"
)

type CondensedView struct {
	Rises  string
	Sets   string
	Length string
	Change string
}

func (c CondensedView) FormatString() string {
	return "" +
		fmt.Sprintf("Rises:  %s\n", c.Rises) +
		fmt.Sprintf("Sets:   %s\n", c.Sets) +
		fmt.Sprintf("Length: %s\n", c.Length) +
		fmt.Sprintf("Change: %s\n", c.Change)
}

func Condensed(query DaylightQuery) CondensedView {
	location := LatLong{
		Lat: query.Lat,
		Lng: query.Long,
	}

	today := SunTimesForPlaceDate(
		location,
		query.Date,
	)

	yesterday := SunTimesForPlaceDate(
		location,
		query.Date.AddDate(0, 0, -1),
	)

	return CondensedView{
		Rises:  LocalisedTime(today.Rises, &query.TZ),
		Sets:   LocalisedTime(today.Sets, &query.TZ),
		Length: FormatDayLength(today),
		Change: FormatLengthDiff(today, yesterday),
	}
}
