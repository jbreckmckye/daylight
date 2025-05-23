package internal

import (
	"errors"
	"time"
)

// Config is a helper that constructs configuration from multiple sources. All fields are nullable.
type Config struct {
	Latitude  *float64
	Longitude *float64
	Timezone  *time.Location
	ForDate   *time.Time
	Condensed *bool
	Json      *bool
	IP        *string
}

// MissingFields allows us to check whether we need more data, e.g. whether we can need API calls.
func (cfg *Config) MissingFields() error {
	if cfg.Latitude == nil {
		return errors.New("missing Latitude data")
	}
	if cfg.Longitude == nil {
		return errors.New("missing Longitude data")
	}
	if cfg.Timezone == nil {
		return errors.New("missing Timezone data")
	}
	// Others (forDate, condensed, json, ip) can be defaulted

	return nil
}

// DaylightQuery turns a Config with nullable values into a DaylightQuery with non-nullables.
// Call MissingFields() first to validate.
func (cfg *Config) DaylightQuery() DaylightQuery {
	// Set defaults
	if cfg.ForDate == nil {
		val := time.Now()
		cfg.ForDate = &val
	}
	if cfg.Condensed == nil {
		val := false
		cfg.Condensed = &val
	}
	if cfg.Json == nil {
		val := false
		cfg.Json = &val
	}
	if cfg.IP == nil {
		val := "n/a"
		cfg.IP = &val
	}

	// Apply timezone to date
	*cfg.ForDate = cfg.ForDate.In(cfg.Timezone)

	return DaylightQuery{
		Lat:       *cfg.Latitude,
		Long:      *cfg.Longitude,
		TZ:        *cfg.Timezone,
		Date:      *cfg.ForDate,
		IP:        *cfg.IP,
		Condensed: *cfg.Condensed,
		Json:      *cfg.Json,
	}
}

// FillValues will fill in missing values from a passed config (without overriding previously set values)
func (cfg *Config) FillValues(cfg2 Config) {
	if cfg.Latitude == nil {
		cfg.Latitude = cfg2.Latitude
	}
	if cfg.Longitude == nil {
		cfg.Longitude = cfg2.Longitude
	}
	if cfg.Timezone == nil {
		cfg.Timezone = cfg2.Timezone
	}
	if cfg.ForDate == nil {
		cfg.ForDate = cfg2.ForDate
	}
	if cfg.Condensed == nil {
		cfg.Condensed = cfg2.Condensed
	}
	if cfg.Json == nil {
		cfg.Json = cfg2.Json
	}
	if cfg.IP == nil {
		cfg.IP = cfg2.IP
	}
}
