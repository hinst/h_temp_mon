package h_temp_mon

import (
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

type TTempReader struct {
	Input  chan uint64
	Output chan float32
	Waiter sync.WaitGroup
}

func NewTempReader() *TTempReader {
	return &TTempReader{}
}

func (this *TTempReader) Read() float32 {
	var data, result = exec.Command("/opt/vc/bin/vcgencmd", "measure_temp").Output()
	var temperature float32 = 0
	if result == nil {
		var text = string(data)
		temperature = ExtractTemperatureFromText(text)
	} else {
		panic(result)
	}
	return temperature
}

func (this *TTempReader) Start() {
	this.Waiter.Add(1)
	this.Output = make(chan float32, 1)
	go this.Run()
}

func (this *TTempReader) Run() {
	for range this.Input {
		var temperature = this.Read()
		this.Output <- temperature
	}
	close(this.Output)
	this.Waiter.Done()
}

func (this *TTempReader) WaitFor() {
	this.Waiter.Wait()
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

func ExtractTemperatureFromText(text string) float32 {
	text = ExtractTemperatureNumberText(text)
	var number, result = strconv.ParseFloat(text, 32)
	if result != nil {
		number = 0
	}
	return float32(number)
}
