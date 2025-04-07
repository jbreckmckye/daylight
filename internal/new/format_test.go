package new

import (
	"testing"
	"time"
)

func TestFormatDayLength(t *testing.T) {
	type TestCase struct {
		SunTimes
		Want string
	}

	cases := []TestCase{
		{
			SunTimes: SunTimes{
				PolarDay:   true,
				PolarNight: false,
				Length:     0,
			},
			Want: "all day (polar sun)",
		},
		{
			SunTimes: SunTimes{
				PolarDay:   false,
				PolarNight: true,
				Length:     0,
			},
			Want: "no time (polar night)",
		},
		{
			SunTimes: SunTimes{
				PolarDay:   false,
				PolarNight: false,
				Length:     (12 * time.Hour) + (34 * time.Minute) + (56 * time.Second),
			},
			Want: "12 hrs, 34 mins",
		},
		{
			SunTimes: SunTimes{
				PolarDay:   false,
				PolarNight: false,
				Length:     24 * time.Hour,
			},
			Want: "24 hrs, 0 mins",
		},
	}

	for _, tc := range cases {
		got := FormatDayLength(tc.SunTimes)
		if got != tc.Want {
			t.Fatalf("expected %q but got %q, SunTimes=%v", tc.Want, got, tc.SunTimes)
		}
	}
}

func TestFormatLengthDiff(t *testing.T) {
	type TestCase struct {
		today     SunTimes
		yesterday SunTimes
		want      string
	}

	cases := []TestCase{
		{
			today:     SunTimes{Length: 12 * time.Hour},
			yesterday: SunTimes{Length: 12 * time.Hour},
			want:      "the same",
		},
		{
			today:     SunTimes{Length: (10 * time.Hour)},
			yesterday: SunTimes{Length: (10 * time.Hour) + (3 * time.Second)},
			want:      "-0m 3s",
		},
		{
			today:     SunTimes{Length: (10 * time.Hour)},
			yesterday: SunTimes{Length: (10 * time.Hour) + (3 * time.Minute) + (3 * time.Second)},
			want:      "-3m 3s",
		},
		{
			today:     SunTimes{Length: (10 * time.Hour) + (3 * time.Second)},
			yesterday: SunTimes{Length: (10 * time.Hour)},
			want:      "+0m 3s",
		},
		{
			today:     SunTimes{Length: (10 * time.Hour) + (3 * time.Minute) + (3 * time.Second)},
			yesterday: SunTimes{Length: (10 * time.Hour)},
			want:      "+3m 3s",
		},
		{ // Extreme case, should never happen
			today:     SunTimes{Length: (11 * time.Hour)},
			yesterday: SunTimes{Length: 10 * time.Hour},
			want:      "+60m 0s",
		},
	}

	for _, tc := range cases {
		got := FormatLengthDiff(tc.today, tc.yesterday)
		if got != tc.want {
			t.Fatalf("expected %q but got %q", got, tc.want)
		}
	}
}
