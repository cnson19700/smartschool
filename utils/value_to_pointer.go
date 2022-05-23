package utils

import (
	"time"
)

func StringP(v string) *string {
	return &v
}

func NumberToPointer(v int) *int {
	return &v
}

func BoolP(v bool) *bool {
	return &v
}

func IntValOrFallback(v *int, fallback int) int {
	if v == nil {
		return fallback
	}
	return *v
}

func DateTimeToPointer(v time.Time) *time.Time {
	return &v
}

func FormatDateTimeP(v time.Time) *time.Time {
	nowUTC := v.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := nowUTC.In(jst)
	return &nowJST
}

func FormatDateTime(v time.Time) time.Time {
	nowUTC := v.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := nowUTC.In(jst)
	return nowJST
}

func SameDate(time1 *time.Time, time2 *time.Time) bool {
	y1, m1, d1 := time1.Date()
	y2, m2, d2 := time2.Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}
