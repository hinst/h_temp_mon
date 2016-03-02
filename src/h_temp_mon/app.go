package h_temp_mon

type TApp struct {
	Ticker     *TTicker
	TempReader *TTempReader
	TempWriter *TTempWriter
	TempDB     *TTempDB
}

func NewApp() *TApp {
	return &TApp{}
}

func (this *TApp) Run() {
	InitializeLog()
	this.TempDB = CreateTempDB()
	this.TempDB.Prepare()
	this.TempDB.Open()
	if this.TempDB.DB == nil {
		return
	}
	this.TempDB.PrepareTables()
	this.Ticker = CreateTicker()
	this.Ticker.Start()
	this.TempReader = CreateTempReader()
	this.TempReader.Input = this.Ticker.Output
	this.TempReader.Start()
	this.TempWriter = CreateTempWriter()
	this.TempWriter.Input = this.TempReader.Output
	this.TempWriter.Start()

	InstallShutdownReceiver(this.Stop)

	this.Ticker.WaitFor()
	this.TempReader.WaitFor()
	this.TempWriter.WaitFor()
	this.TempDB.Close()
}

func (this *TApp) Stop() {
	this.Ticker.Stop()
}
