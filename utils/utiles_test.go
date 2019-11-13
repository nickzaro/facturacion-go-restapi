package utils

import (
	"testing"
)

func TestConvertirAnioMesString(t *testing.T) {
	date := "2019-05-16T00:00:00"
	if ConvertirAnioMesString(date) != "2019-05" {
		t.Error()
	}
}

func TestConvertirAPesos(t *testing.T) {
	if ConvertirAPesos(100, "USD") != 6500 {
		t.Error()
	}
	if ConvertirAPesos(100, "USX") != 100 {
		t.Error()
	}
}
