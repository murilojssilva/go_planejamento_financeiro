package util

import "testing"

func TestStringToDate(t *testing.T) {
	var convertedTime = StringToTime("2023-07-12T15:30:26")

	if convertedTime.Year() != 2023 {
		t.Errorf("Espera que o ano seja 2023")
	}
}