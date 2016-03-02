package h_temp_mon

import (
	"sync"
	"time"

	_ "github.com/nakagami/firebirdsql"
)

type TTempWriter struct {
	Input  chan float32
	Waiter sync.WaitGroup
	DB     *TTempDB
}

func CreateTempWriter() *TTempWriter {
	var this = &TTempWriter{}
	return this
}

func (this *TTempWriter) Start() {
	this.Waiter.Add(1)
	go func() {
		defer this.Waiter.Done()
		this.Run()
	}()
}

func (this *TTempWriter) Run() {
	for temperature := range this.Input {
		Log.Println(TimeToFirebirdString(time.Now()))
		this.Write(temperature)
	}
}

func (this *TTempWriter) WaitFor() {
	this.Waiter.Wait()
}

func (this *TTempWriter) Write(temperature float32) {
	this.DB.WriteTemperature(temperature)
}
