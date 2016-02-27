package h_temp_mon

import (
	"os/exec"
	"strings"
)

type TTempReader struct {
	Input chan int64
}

func (this *TTempReader) Read() {
	data, result := exec.Command("/opt/vc/bin/vcgencmd measure_temp").Output()
	if result == nil {
		var text = string(data)
		unused(text)
	}
}

func ExtractTemperatureNumberText(text string) string {
	const desiredChars = "0123456789."
	var result = ""
	var inside = false
	for _, character := range text {
		if strings.IndexRune(desiredChars, character) != -1 {
			if false == inside {
				inside = true
			}
			result = result + string(character)
		} else if inside {
			break
		}
	}
	return result
}
