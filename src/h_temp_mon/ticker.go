package h_temp_mon

import "time"

type TTicker struct {
	Interval   time.Duration
	BufferSize int
	Counter    uint64
	ShouldStop bool
	Output     chan uint64
}

func NewTicker() *TTicker {
	var result = &TTicker{}
	result.BufferSize = 1
	result.Interval = time.Second
	result.Counter = 0
	return result
}

func (this *TTicker) Start() {
	this.Output = make(chan uint64, this.BufferSize)
}

func (this *TTicker) Run() {
	for {
		this.Counter++
		this.Output <- this.Counter
		time.Sleep(this.Interval)
		if this.ShouldStop {
			break
		}
	}
	close(this.Output)
}

func (this *TTicker) Stop() {
	this.ShouldStop = true
}
