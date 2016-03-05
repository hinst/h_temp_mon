package h_temp_mon

import (
	"os"
	"path/filepath"
	"time"
)

type TApp struct {
	AppDirectory string
	AppURL       string
	Ticker       *TTicker
	TempReader   *TTempReader
	TempWriter   *TTempWriter
	TempDB       *TTempDB
	WebUI        *TWebUI
}

func NewApp() *TApp {
	return &TApp{}
}

func (this *TApp) Run() {
	this.AppDirectory, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	this.AppURL = "/h_temp_mon"
	InitializeLog()
	Log.Println("Now running; AppDirectory='" + this.AppDirectory + "'")
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
	this.WebUI = CreateWebUI()
	this.WebUI.URL = this.AppURL
	this.WebUI.Prepare()

	InstallShutdownReceiver(this.Stop)

	this.Ticker.WaitFor()
	this.TempReader.WaitFor()
	this.TempWriter.WaitFor()
	this.TempDB.Close()
}

func (this *TApp) Stop() {
	this.Ticker.Stop()
}
