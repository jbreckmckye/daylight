package cmd

import (
	"fmt"
	"time"

	"github.com/alexflint/go-arg"

	daylight "github.com/jbreckmckye/daylight/internal"
	"github.com/jbreckmckye/daylight/internal/api"
)

type ExitCode = int

const (
	exitOK  ExitCode = 0
	exitErr ExitCode = 1
)

type DaylightQuery struct {
	Lat  float64
	Long float64
	TZ   time.Location
	Date time.Time
}

type arguments struct {
	Short     bool    `help:"Show in condensed format"`
	Latitude  *float64 `help:"Set latitude (requires --longitude)"`
	Longitude *float64 `help:"Set longitude (requires --latitude)"`
	Date      string   `help:"Date in YYYY-MM-DD"`
	Timezone  string   `help:"Timezone in IANA format e.g. 'Europe/London'"`
}

type inputParse struct {
	lat   *float64
	long  *float64
	tz    *time.Location
	date  time.Time
	short bool
	ip    string
}

func (parsed *inputParse) daylightQuery() DaylightQuery {
	return DaylightQuery{
		Lat: *parsed.lat,
		Long: *parsed.long,
		TZ: *parsed.tz,
		Date: parsed.date,
	}
}

func (parsed *inputParse) readFromArgs(args *arguments) error {
	if args.Latitude != nil {
		parsed.lat = args.Latitude
	}

	if args.Longitude != nil {
		parsed.long = args.Longitude
	}

	if args.Timezone != "" {
		timezone, err := time.LoadLocation(args.Timezone)
		if err != nil {
			return fmt.Errorf("parsing timezone: %w", err)
		}
		parsed.tz = timezone
	}

	if args.Date != "" {
		date, err := time.Parse(time.DateOnly, args.Date)
		if err != nil {
			return fmt.Errorf("parsing date: %w", err)
		}
		parsed.date = date
	}

	parsed.short = args.Short

	return nil
}

func (parsed *inputParse) setDefaults() {
	parsed.date = time.Now()
}

func (parsed *inputParse) hasEnoughData() bool {
	return parsed.lat != nil && parsed.long != nil && parsed.tz != nil
}

func (parsed *inputParse) readFromAPI() error {
	ipinfo, err := api.FetchIPInfo()
	if err != nil {
		return err
	}

	parsed.ip = ipinfo.IP

	latlong, err := daylight.LocationToLatLong(ipinfo.Loc)
	if err != nil {
		return err
	}

  if parsed.lat == nil {
		parsed.lat = &latlong.Lat
	}

	if parsed.long == nil {
		parsed.long = &latlong.Lng
	}

	if parsed.tz == nil {
		timezone, err := time.LoadLocation(ipinfo.TZ)
		if err != nil {
			return err
		}
		parsed.tz = timezone
	}

	return nil
}

func Daylight() ExitCode {
	var args arguments
  var printErr = func (err error) {
		fmt.Printf("[daylight] %s\n", err)
	}

	err := parseArgs(&args)
	if err != nil {
		printErr(err)
		return exitErr
	}

	var input = inputParse{}
	input.setDefaults()

  err = input.readFromArgs(&args)
	if err != nil {
		printErr(err)
		return exitErr
	}

	if !input.hasEnoughData() {
		fmt.Println("[debug] load from ipinfo")
		err = input.readFromAPI()
    if err != nil {
			printErr(err)
			return exitErr
		}
	}

	query := input.daylightQuery()
	fmt.Printf("[debug] lat is %v\n", query.Lat)
	fmt.Printf("[debug] long is %v\n", query.Long)
	fmt.Printf("[debug] date is %v\n", query.Date)
	fmt.Printf("[debug] tz is %v\n", query.TZ.String())

	fmt.Printf("short mode? %t\n", args.Short)

	return exitOK
}

func parseArgs(args *arguments) error {
	arg.MustParse(args)

	if (args.Latitude == nil) != (args.Longitude == nil) {
		return fmt.Errorf("--latitude and --longitude must both be set, if used")
	}

	if (args.Latitude != nil) && ((*args.Latitude < -90) || (*args.Latitude > 90)) {
		return fmt.Errorf("--latitude must be between -90 and 90")
	}

	if (args.Longitude != nil) && ((*args.Longitude < -180) || (*args.Longitude > 180)) {
		return fmt.Errorf("--longitude must be between -180 and 180")
	}

	return nil
}
