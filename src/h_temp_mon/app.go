package h_temp_mon

type TApp struct {
	Ticker     *TTicker
	TempReader *TTempReader
	TempWriter *TTempWriter
}

func NewApp() *TApp {
	return &TApp{}
}

func (this *TApp) Run() {
	this.Ticker = CreateTicker()
	this.Ticker.Start()
	this.TempReader = CreateTempReader()
	this.TempReader.Input = this.Ticker.Output
	this.TempReader.Start()
	this.TempWriter = CreateTempWriter()
	this.TempWriter.Prepare()

	InstallShutdownReceiver(this.Stop)

	this.Ticker.WaitFor()
	this.TempReader.WaitFor()
}

func (this *TApp) Stop() {
	this.Ticker.Stop()
}
