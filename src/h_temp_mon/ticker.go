package h_temp_mon

import (
	"sync"
	"time"
)

type TTicker struct {
	Interval   time.Duration
	BufferSize int
	Counter    uint64
	ShouldStop bool
	Output     chan uint64
	Waiter     sync.WaitGroup
}

func CreateTicker() *TTicker {
	var result = &TTicker{}
	result.BufferSize = 0
	result.Interval = time.Second
	result.Counter = 0
	return result
}

func (this *TTicker) Start() {
	this.Waiter.Add(1)
	this.Output = make(chan uint64, this.BufferSize)
	go func() {
		defer this.Waiter.Done()
		defer close(this.Output)
		this.Run()
	}()
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
}

func (this *TTicker) Stop() {
	this.ShouldStop = true
}

func (this *TTicker) WaitFor() {
	this.Waiter.Wait()
}
