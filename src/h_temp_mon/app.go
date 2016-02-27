package h_temp_mon

import "fmt"

type TApp struct {
	Ticker     *TTicker
	TempReader *TTempReader
}

func NewApp() *TApp {
	return &TApp{}
}

func (this *TApp) Run() {
	this.Ticker = NewTicker()
	this.Ticker.Start()
	this.TempReader = NewTempReader()
	this.TempReader.Input = this.Ticker.Output
	this.TempReader.Start()
	InstallShutdownReceiver(this.Stop)
	for value := range this.TempReader.Output {
		fmt.Println(value)
	}
	this.Ticker.WaitFor()
	this.TempReader.WaitFor()
}

func (this *TApp) Stop() {
	this.Ticker.Stop()
}
