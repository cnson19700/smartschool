package apptypes

import (
	"encoding/json"
	"time"

	"github.com/smartschool/utils"
)

type TimeNanoSec struct {
	Value time.Time
}

func NewTimeNanoSec(v *time.Time) *TimeNanoSec {
	return &TimeNanoSec{Value: *v}
}

func (t *TimeNanoSec) TimeP() *time.Time {
	if t == nil {
		return nil
	}
	return &t.Value
}

func (t *TimeNanoSec) UnixTime() int64 {
	if t == nil {
		return 0
	}
	return t.Value.Unix()
}

func (t *TimeNanoSec) Add(value int) *TimeNanoSec {
	time := t.Value.Add(time.Duration(value) * time.Second)
	newTime := TimeNanoSec{
		Value: time,
	}
	return &newTime
}

func (t *TimeNanoSec) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Value)
}

func (t *TimeNanoSec) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	value, _ := utils.ParseTimeFromRFC3339(v)
	t.Value = value
	return nil
}
