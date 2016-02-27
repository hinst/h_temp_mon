package h_temp_mon

import "testing"

func testExtractTemperatureNumberText(source string, expectedResult string) {
	var matched = ExtractTemperatureNumberText(source) == expectedResult
	if false == matched {
		panic("mismatch")
	}
}

func TestExtractTemperatureNumberText(t *testing.T) {
	testExtractTemperatureNumberText("asdf=67.8'degreesOfC", "67.8")
}
