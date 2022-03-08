package helper

import "time"

func StringToTimeUTC(stringTime string) time.Time {

	t, err := time.Parse(time.RFC3339, stringTime)

	if err != nil {
		panic(err)
	}

	return t.UTC()
}
