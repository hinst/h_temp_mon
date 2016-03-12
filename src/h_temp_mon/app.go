package h_temp_mon

import (
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type TApp struct {
	Directory  string
	URL        string
	Ticker     *TTicker
	TempReader *TTempReader
	TempWriter *TTempWriter
	TempDB     *TTempDB
	WebUI      *TWebUI
}

func NewApp() *TApp {
	return &TApp{}
}

func (this *TApp) Run() {
	this.Directory, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	this.URL = "/h_temp_mon"
	InitializeLog()
	Log.Println("Now running; AppDirectory='" + this.Directory + "'")
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
	this.WebUI.Directory = this.Directory
	this.WebUI.URL = this.URL
	this.WebUI.DB = this.TempDB
	this.WebUI.Prepare()
	go http.ListenAndServe(":9001", nil)

	InstallShutdownReceiver(this.Stop)

	this.Ticker.WaitFor()
	this.TempReader.WaitFor()
	this.TempWriter.WaitFor()
	this.TempDB.Close()
}

func (this *TApp) Stop() {
	this.Ticker.Stop()
}
