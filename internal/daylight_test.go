package internal

import (
	"testing"
	"time"
)

func TestLocationToLatLong(t *testing.T) {
	cases := []struct {
		name  string
		input string
		ok    bool
		lat   float64
		lng   float64
	}{
		{
			name:  "empty string",
			input: "",
			lat:   0,
			lng:   0,
			ok:    false,
		},
		{
			name:  "invalid string, not a lat long",
			input: "undefined",
			ok:    false,
			lat:   0,
			lng:   0,
		},
		{
			name:  "invalid string, degree format",
			input: "40° 26.767′ N 79° 58.933′ W",
			ok:    false,
			lat:   0,
			lng:   0,
		},
		{
			name:  "positive,positive",
			input: "42.7865,15.9810",
			ok:    true,
			lat:   42.7865,
			lng:   15.9810,
		},
		{
			name:  "positive,negative",
			input: "42.7865,-5.9810",
			ok:    true,
			lat:   42.7865,
			lng:   -5.9810,
		},
		{
			name:  "negative,positive",
			input: "-89.7865,15.9810",
			ok:    true,
			lat:   -89.7865,
			lng:   15.9810,
		},
		{
			name:  "negative,negative",
			input: "-1.7865,-0.9810",
			ok:    true,
			lat:   -1.7865,
			lng:   -0.9810,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, gotErr := LocationToLatLong(tc.input)

			if !tc.ok {
				if gotErr == nil {
					t.Fatalf("input %q; expected an error but got nil", tc.input)
				}
				return
			} else {
				if gotErr != nil {
					t.Fatalf("input %q; expected no errors but got error %q", tc.input, gotErr.Error())
				}
				if got.Lat != tc.lat {
					t.Fatalf("input %q; expected latitude %f but got %f", tc.input, tc.lat, got.Lat)
				}
				if got.Lng != tc.lng {
					t.Fatalf("input %q; expected longitude %f but got %f", tc.input, tc.lng, got.Lng)
				}
			}
		})
	}
}

func TestSunTimesForPlaceDate(t *testing.T) {
	cases := []struct {
		name         string
		inputLatLong LatLong
		inputDate    time.Time
		rises        time.Time
		sets         time.Time
		polarNight   bool
		polarDay     bool
	}{
		{
			name: "Svalbard, Norway, in winter - polar night",
			inputLatLong: LatLong{
				Lat: 77.8750,
				Lng: 20.9752,
			},
			inputDate:  utcTime("2025-02-03T10:00:00+01:00"),
			rises:      time.Time{}, // expect empty
			sets:       time.Time{}, // expect empty
			polarNight: true,
			polarDay:   false,
		},
		{
			name: "South pole, in winter - polar day",
			inputLatLong: LatLong{
				Lat: -90,
				Lng: 0,
			},
			inputDate:  utcTime("2025-02-03T10:00:00Z"),
			rises:      time.Time{}, // expect empty
			sets:       time.Time{}, // expect empty
			polarNight: false,
			polarDay:   true,
		},
		{
			name: "Cape Town, in winter - long day",
			inputLatLong: LatLong{
				Lat: -33.9258,
				Lng: 18.4232,
			},
			inputDate:  utcTime("2025-02-03T10:00:00+02:00"),
			rises:      utcTime("2025-02-03T06:09:56+02:00"), // SAA time is +2 hours from UTC
			sets:       utcTime("2025-02-03T19:50:17+02:00"),
			polarNight: false,
			polarDay:   false,
		},
		{
			name: "London, in winter - short day",
			inputLatLong: LatLong{
				Lat: 51.5072,
				Lng: 0.1276,
			},
			inputDate:  utcTime("2025-02-04T10:00:00Z"),
			rises:      utcTime("2025-02-04T07:33:00Z"),
			sets:       utcTime("2025-02-04T16:53:49Z"),
			polarNight: false,
			polarDay:   false,
		},
		{
			name: "London, in summer - long day",
			inputLatLong: LatLong{
				Lat: 51.5072,
				Lng: 0.1276,
			},
			inputDate:  utcTime("2025-06-20T10:00:00+01:00"),
			rises:      utcTime("2025-06-20T04:41:55+01:00"),
			sets:       utcTime("2025-06-20T21:20:12+01:00"),
			polarNight: false,
			polarDay:   false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			suntimes := SunTimesForPlaceDate(tc.inputLatLong, tc.inputDate)

			if tc.polarNight && !suntimes.PolarNight {
				t.Fatal("expected a polar night, but it wasn't")
			}
			if tc.polarDay && !suntimes.PolarDay {
				t.Fatal("expected a polar day, but it wasn't")
			}
			if tc.rises != suntimes.Rises {
				t.Fatalf("expected sunrise at %v; got %v", tc.rises, suntimes.Rises)
			}
			if tc.sets != suntimes.Sets {
				t.Fatalf("expected sunset at %v; got %v", tc.sets, suntimes.Sets)
			}
		})
	}
}

func TestApproximateNoon(t *testing.T) {
	s := SunTimes{
		Rises: utcTime("2025-06-20T07:00:00Z"),
    Length: (time.Hour * 14) + (time.Minute * 40),
	}

	want := utcTime("2025-06-20T14:20:00Z")
	got := s.ApproximateNoon()
	
	if got != want {
		t.Fatalf("wanted %v, got %v, for %v", want, got, s)
	}
}

func utcTime(s string) time.Time {
	result, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return result.UTC()
}
