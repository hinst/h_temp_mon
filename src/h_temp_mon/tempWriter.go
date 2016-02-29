package h_temp_mon

import (
	"database/sql"
	"io/ioutil"
	"sync"

	_ "github.com/nakagami/firebirdsql"
)

type TTempWriter struct {
	DBUserName string
	DBPassword string
	DBName     string
	DB         *sql.DB
	Input      chan float32
	Waiter     sync.WaitGroup
}

func CreateTempWriter() *TTempWriter {
	var this = &TTempWriter{}
	this.DBUserName = "h_temp_mon"
	this.DBPassword = "password"
	this.DBName = "h_temp_mon"
	return this
}

func (this *TTempWriter) Prepare() {
	var data, readResult = ioutil.ReadFile("dbpassword.txt")
	if readResult == nil {
		this.DBPassword = string(data)
	}
}

func (this *TTempWriter) OpenDB() {
	var connectionString = this.GetConnectionString()
	var dbOpenResult error
	Log.Println("OpenDB: DBName=" + this.DBName)
	this.DB, dbOpenResult = sql.Open("firebirdsql", connectionString)
	if dbOpenResult == nil {
		Log.Println("OpenDB: success")
	} else {
		Log.Println("OpenDB: fail: ", dbOpenResult)
		this.DB = nil
	}
}

func (this *TTempWriter) CloseDB() {
	if this.DB != nil {
		this.DB.Close()
	}
}

func (this *TTempWriter) Start() {
	this.Waiter.Add(1)
	go func() {
		defer this.Waiter.Done()
		this.Run()
	}()
}

func (this *TTempWriter) Run() {
	this.OpenDB()
	for temperature := range this.Input {
		unused(temperature)
	}
	defer this.CloseDB()
}

func (this *TTempWriter) GetConnectionString() string {
	return this.DBUserName + ":" + this.DBPassword + "@localhost/" + this.DBName
}

func (this *TTempWriter) WaitFor() {
	this.Waiter.Wait()
}
