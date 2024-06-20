package comm_utils

import (
	"strconv"
	"time"
)

var ZeroTime = time.Time{}
var StartOfTime, _ = ParseRFC3339Time("1970-01-01T00:00:00+00:00")
var EndOfTime, _ = ParseRFC3339Time("3000-01-01T00:00:00+00:00")
var DefaultLoc, _ = time.LoadLocation("Asia/Singapore")

func ParseRFC3339Time(in string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, in)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func MustParseRFC3339Time(in string) time.Time {
	t, err := time.Parse(time.RFC3339, in)
	if err != nil {
		panic(err)
	}
	return t
}

func MustParseToInt(in string) int {
	out, err := strconv.Atoi(in)
	if err != nil {
		panic(err)
	}
	return out
}

func MustParseToUint(in string) uint {
	out, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		panic(err)
	}
	return uint(out)
}

func BeginOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location()).AddDate(0, 0, 1).Add(-1 * time.Second)
}

func BeginOfWeek(t time.Time) time.Time {
	year, month, day := t.Date()
	t = time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	for t.Weekday() != time.Monday {
		t = t.AddDate(0, 0, -1)
	}

	return t
}

func EndOfWeek(t time.Time) time.Time {
	year, month, day := t.Date()
	t = time.Date(year, month, day, 0, 0, 0, 0, t.Location()).AddDate(0, 0, 1).Add(-1 * time.Second)
	for t.Weekday() != time.Sunday {
		t = t.AddDate(0, 0, 1)
	}

	return t
}

func BeginOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
}

func EndOfMonth(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 1, 0, 0, 0, 0, t.Location()).AddDate(0, 1, 0).Add(-1 * time.Second)
}
