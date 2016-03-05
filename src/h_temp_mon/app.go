package h_temp_mon

import (
	"os"
	"path/filepath"
	"time"
)

type TApp struct {
	Ticker     *TTicker
	TempReader *TTempReader
	TempWriter *TTempWriter
	TempDB     *TTempDB
}

func NewApp() *TApp {
	return &TApp{}
}

var AppDirectory string

func InitializeAppPath() {
	AppDirectory, _ = filepath.Abs(filepath.Dir(os.Args[0]))
}

func (this *TApp) Run() {
	InitializeAppPath()
	InitializeLog()
	Log.Println("Now running; AppDirectory='" + AppDirectory + "'")
	this.TempDB = CreateTempDB()
	this.TempDB.Prepare()
	this.TempDB.Open()
	if this.TempDB.DB == nil {
		return
	}
	this.TempDB.PrepareTables()
	this.Ticker = CreateTicker()
	this.Ticker.Interval = time.Second * 3
	this.Ticker.Start()
	this.TempReader = CreateTempReader()
	this.TempReader.Input = this.Ticker.Output
	this.TempReader.Start()
	this.TempWriter = CreateTempWriter()
	this.TempWriter.Input = this.TempReader.Output
	this.TempWriter.DB = this.TempDB
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
