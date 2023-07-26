package util

import "testing"

func TestStringToDate(t *testing.T) {
	var convertedTime = StringToTime("2023-07-26T10:10:10")

	if convertedTime.Year() != 2023 {
		t.Errorf("Espera que o ano seja 2023")
	}

	if convertedTime.Month() != 07 {
		t.Errorf("Espera que o mês seja 7")
	}

	if convertedTime.Day() != 26 {
		t.Errorf("Espera que o dia seja 26")
	}

	if convertedTime.Hour() != 10 {
		t.Errorf("Espera que a hora sejam 10")
	}

	if convertedTime.Minute() != 10 {
		t.Errorf("Espera que os minutos sejam 10")
	}

	if convertedTime.Second() != 10 {
		t.Errorf("Espera que os segundos sejam 10")
	}
}
