package h_temp_mon

import (
	"sync"
	"time"
)

type TTicker struct {
	AtomicInterval time.Duration
	Interval       time.Duration
	BufferSize     int
	Counter        uint64
	ShouldStop     bool
	Output         chan uint64
	Waiter         sync.WaitGroup
}

func CreateTicker() *TTicker {
	var result = &TTicker{}
	result.BufferSize = 0
	result.AtomicInterval = time.Second / 2
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
	var currentInterval time.Duration = 0
	for {
		this.Counter++
		if this.Interval < currentInterval {
			this.Output <- this.Counter
			currentInterval -= this.Interval
		}
		time.Sleep(this.AtomicInterval)
		currentInterval += this.AtomicInterval
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
