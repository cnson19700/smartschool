package helper

import (
	"fmt"
	"time"
)

func GetMonthSelectBox() string {
	result := `<select class="form-control" name="month" id="month">`

	currentMonth := int(time.Now().Month())
	for i := 1; i <= 12; i++ {
		if i == currentMonth {
			result += fmt.Sprintf(`<option value="%v" selected>%v</option>`, i, i)
		} else {
			result += fmt.Sprintf(`<option value="%v">%v</option>`, i, i)
		}
	}

	result += `</select>`

	return result
}
