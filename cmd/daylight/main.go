package main

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"time"
	_ "time/tzdata"

	"github.com/alexflint/go-arg"

	daylight "github.com/jbreckmckye/daylight/internal"
	templates "github.com/jbreckmckye/daylight/internal/templates"
)

func main() {
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

	loc := first(args.Loc, ipInfo.Loc)
	// If loc was supplied as an arg it might have an escaped negative
	loc = strings.Replace(loc, "\\", "", -1)

	latlong, err := daylight.LocationToLatLong(loc)
	checkErr(err)

	tz := first(args.Timezone, ipInfo.TZ)
	timezone, err := time.LoadLocation(tz)
	checkErr(err)

	now := time.Now().In(timezone)
	if args.Date != "" {
		now, err = time.Parse(time.DateOnly, args.Date)
		checkErr(err)
	}

	viewmodel := daylight.TodayStats(now, timezone, latlong, ipInfo.IP)

	tmpl := templates.TodayTemplate()
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, viewmodel)
	checkErr(err)

	var output string

	if args.Short {
		output = fmt.Sprintf(
			"Rises:  %s\nSets:   %s\nLength: %s\nChange:  %s",
			viewmodel.Rise, viewmodel.Sets, viewmodel.Len, viewmodel.Diff,
		)
	} else if daylight.UsePrettyMode() {
		output = daylight.Sunnify(buf.String())
	} else {
		output = buf.String()
	}

	fmt.Println(output)

	renders := render(viewmodel)
	fmt.Println(renders)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func first(strings ...string) string {
	for _, s := range strings {
		if s != "" {
			return s
		}
	}
	return ""
}
