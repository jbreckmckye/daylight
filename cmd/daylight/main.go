package main

import (
	"fmt"
	"log"
	"strings"
	"time"
	_ "time/tzdata"

	"github.com/alexflint/go-arg"

	daylight "github.com/jbreckmckye/daylight/internal"
)

func main2() {
	log.SetPrefix("[daylength] ")
	log.SetFlags(0)

	var args struct {
		Short    bool   `help:"Show in condensed format"`
		Loc      string `help:"Set latitude, longitude in format 'NN.nn,NN.nn'"`
		Date     string `help:"Date in YYYY-MM-DD"`
		Timezone string `help:"Timezone e.g. 'Europe/London'"`
	}

	arg.MustParse(&args)

	ipInfo, err := daylight.FetchIPInfo()
	checkErr(err)

	loc := firstNonEmpty(args.Loc, ipInfo.Loc)
	// If loc was supplied as an arg it might have an escaped negative
	loc = strings.Replace(loc, "\\", "", -1)

	latlong, err := daylight.LocationToLatLong(loc)
	checkErr(err)

	tz := firstNonEmpty(args.Timezone, ipInfo.TZ)
	timezone, err := time.LoadLocation(tz)
	checkErr(err)

	now := time.Now().In(timezone)
	if args.Date != "" {
		now, err = time.Parse(time.DateOnly, args.Date)
		checkErr(err)
	}

	todayView := daylight.TodayStats(now, timezone, latlong, ipInfo.IP)

	if args.Short {
		fmt.Printf(
			"Rises:  %s\nSets:   %s\nLength: %s\nChange:  %s\n",
			todayView.Rise, todayView.Sets, todayView.Len, todayView.Diff,
		)
	} else {
		renders := render(todayView)
		fmt.Println(renders)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func firstNonEmpty(strings ...string) string {
	for _, s := range strings {
		if s != "" {
			return s
		}
	}
	return ""
}
