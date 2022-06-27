package helper

import "time"

func StringToTimeUTC(stringTime string) (time.Time, error) {

	t, err := time.Parse(time.RFC3339, stringTime)

	return t.UTC(), err
}

func StringToDateUTC(stringTime string) (time.Time, error) {

	t, err := time.Parse("02/01/2006", stringTime)

	return t, err
}
