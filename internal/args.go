package internal

import (
	"fmt"
	"time"

	"github.com/alexflint/go-arg"
)

// Arguments is the set of arg flags we pass to go-arg
type Arguments struct {
	Latitude  *float64 `help:"Set latitude (requires --longitude)"`
	Longitude *float64 `help:"Set longitude (requires --latitude)"`
	Timezone  *string  `help:"Timezone in IANA format e.g. 'Europe/London'"`
	Date      *string  `help:"Date in YYYY-MM-DD"`
	Short     *bool    `help:"Show in condensed format"`
	Json      *bool    `help:"Short JSON output"`
	Version   *bool    `help:"Show version"`
}

// ReadFromCLI reads command line flags into Arguments
func (args *Arguments) ReadFromCLI() {
	arg.MustParse(args)
}

// Config turns an Arguments obj into an internal.Config, performing validation
func (args *Arguments) Config() (cfg Config, err error) {
	if (args.Latitude == nil) != (args.Longitude == nil) {
		return cfg, fmt.Errorf("--latitude and --longitude must both be set, if used")
	}

	if (args.Latitude != nil) && ((*args.Latitude < -90) || (*args.Latitude > 90)) {
		return cfg, fmt.Errorf("--latitude must be between -90 and 90")
	}

	if (args.Longitude != nil) && ((*args.Longitude < -180) || (*args.Longitude > 180)) {
		return cfg, fmt.Errorf("--longitude must be between -180 and 180")
	}

	cfg.Latitude = args.Latitude
	cfg.Longitude = args.Longitude

	if args.Timezone != nil {
		cfg.Timezone, err = time.LoadLocation(*args.Timezone)
		if err != nil {
			return cfg, fmt.Errorf("--timezone was not found")
		}
	}

	if args.Date != nil {
		date, err := time.Parse(time.DateOnly, *args.Date)
		if err != nil {
			return cfg, fmt.Errorf("--date was not a valid date")
		}
		cfg.ForDate = &date
	}

	if args.Short != nil {
		cfg.Condensed = args.Short
	}

	if args.Json != nil {
		cfg.Json = args.Json
	}

	return cfg, nil
}
