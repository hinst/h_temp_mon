package h_temp_mon

import "fmt"

type TApp struct {
	Ticker     TTicker
	TempReader TTempReader
}

func NewApp() *TApp {
	return &TApp{}
}

func (this *TApp) Run() {
	this.Ticker.Start()
	InstallShutdownReceiver(this.Stop)
	this.Ticker.Waiter.Wait()
}

func (this *TApp) Stop() {
	fmt.Println("stop")
	this.Ticker.Stop()
}
