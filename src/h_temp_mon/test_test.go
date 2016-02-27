package h_temp_mon

import "testing"

func TestExtractTemperatureNumberText(t *testing.T) {
	var result = ExtractTemperatureNumberText("asdf=67.8'degreesOfC") == "67.9"
	if false == result {
		t.Error()
	}
}
