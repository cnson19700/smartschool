package utils

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/senseyeio/duration"
)

const (
	dateRegexString     = `^(?P<year>[0-9]{4})(?:[\-/]?(?P<month>1[0-2]|0[1-9])(?:[\-/]?(?P<day>3[01]|0[1-9]|[12][0-9]))?)?$`
	dateTimeRegexString = `^(?P<year>[0-9]{4})[\-/]?(?P<month>1[0-2]|0[1-9])[\-/]?(?P<day>3[01]|0[1-9]|[12][0-9])[ T](?P<hour>2[0-3]|[01][0-9]):?(?P<min>[0-5][0-9])(?::?(?P<sec>[0-5][0-9])(?:\.(?P<nsec>\d{1,9}))?)?(?P<loc>Z|[+\-](?:2[0-3]|[01][0-9])(?::?(?:[0-5][0-9]))?)?$`
)

var (
	dateRegex     = regexp.MustCompile(dateRegexString)
	dateTimeRegex = regexp.MustCompile(dateTimeRegexString)
)

func IsDateString(text string) bool {
	return dateRegex.MatchString(text)
}

func IsDateTimeString(text string) bool {
	return dateTimeRegex.MatchString(text)
}

func ParseDateString(text string) (t time.Time, ok bool) {
	match := dateRegex.FindStringSubmatch(text)
	if match == nil {
		return
	}

	year := 1970
	month := time.January
	day := 1

	for i, name := range dateRegex.SubexpNames() {

		if i < 0 || len(match) <= i {
			continue
		}

		s := match[i]
		if len(s) == 0 {
			continue
		}

		switch name {
		case "year":
			if n, err := strconv.ParseInt(s, 10, 0); err == nil {
				year = int(n)
			}
		case "month":
			if n, err := strconv.ParseInt(s, 10, 0); err == nil {
				month = time.Month(n)
			}
		case "day":
			if n, err := strconv.ParseInt(s, 10, 0); err == nil {
				day = int(n)
			}
		}
	}

	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC), true
}

func ParseDateTimeString(text string) (t time.Time, ok bool) {
	match := dateTimeRegex.FindStringSubmatch(text)
	if match == nil {
		return
	}

	year := 1970
	month := time.January
	day := 1
	hour := 0
	min := 0
	sec := 0
	nsec := 0
	loc := time.UTC

	for i, name := range dateTimeRegex.SubexpNames() {

		if i < 0 || len(match) <= i {
			continue
		}

		s := match[i]
		if len(s) == 0 {
			continue
		}

		switch name {
		case "year":
			if n, err := strconv.ParseInt(s, 10, 0); err == nil {
				year = int(n)
			}
		case "month":
			if n, err := strconv.ParseInt(s, 10, 0); err == nil {
				month = time.Month(n)
			}
		case "day":
			if n, err := strconv.ParseInt(s, 10, 0); err == nil {
				day = int(n)
			}
		case "hour":
			if n, err := strconv.ParseInt(s, 10, 0); err == nil {
				hour = int(n)
			}
		case "min":
			if n, err := strconv.ParseInt(s, 10, 0); err == nil {
				min = int(n)
			}
		case "sec":
			if n, err := strconv.ParseInt(s, 10, 0); err == nil {
				sec = int(n)
			}
		case "nsec":
			for n := 9 - len(s); n > 0; n-- {
				s += "0"
			}
			if n, err := strconv.ParseInt(s, 10, 0); err == nil {
				nsec = int(n)
			}
		case "loc":
			if s != "Z" {
				offset := 0
				if n, err := strconv.ParseInt(s[1:3], 10, 0); err == nil {
					offset += int(n) * 60
				}
				if len(s) >= 5 {
					if n, err := strconv.ParseInt(s[len(s)-2:], 10, 0); err == nil {
						offset += int(n)
					}
				}
				if s[0] == '-' {
					offset *= -1
				}
				loc = time.FixedZone(s, offset*60)
			}
		}
	}

	return time.Date(year, month, day, hour, min, sec, nsec, loc), true
}

func DatetimeToStringFormat(t *time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func StringToDatetimeFormat(str string) (date time.Time, err error) {
	layout := "2006-01-02"
	date, err = time.Parse(layout, str)
	return date, err
}

func ParseTimeFromRFC3339OrNow(str string) (time.Time, error) {
	if str == "" {
		return time.Now(), nil
	}
	return time.Parse(time.RFC3339, str)
}

func ParseTimeFromRFC3339(str string) (time.Time, error) {
	return time.Parse(time.RFC3339, str)
}

func GetNowTimestamp() int64 {
	return time.Now().Unix()
}

func GetNowTimestampP() *int {
	now := int(GetNowTimestamp())
	return &now
}

func CompareTime(time1 int64, time2 int64) bool {
	return time1 == time2
}

func ParseISO8601ToSecond(durationStr string) int {
	d, _ := duration.ParseISO8601(durationStr)
	t := (d.D * 24 * 3600) + (d.TH * 3600) + (d.TM * 60) + (d.TS)
	return t
}

func ParseSecondToISO8601(second int) string {
	secondStr := strconv.Itoa(second)
	duration, _ := time.ParseDuration(secondStr + "s")
	return "PT" + strings.ToUpper(duration.String())
}
