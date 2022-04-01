package helper

import "time"

func StringToTimeUTC(stringTime string) (time.Time, error) {

	t, err := time.Parse(time.RFC3339, stringTime)

	return t.UTC(), err
}
