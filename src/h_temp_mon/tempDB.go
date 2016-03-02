package h_temp_mon

import (
	"database/sql"
	"io/ioutil"
)

type TTempDB struct {
	DBUserName string
	DBPassword string
	DBName     string
	DB         *sql.DB
}

func CreateTempDB() *TTempDB {
	this := &TTempDB{}
	this.DBUserName = "h_temp_mon"
	this.DBPassword = "password"
	this.DBName = "h_temp_mon"
	return this
}

func (this *TTempDB) Prepare() {
	var data, readResult = ioutil.ReadFile("dbpassword.txt")
	if readResult == nil {
		this.DBPassword = string(data)
	}
}

func (this *TTempDB) GetConnectionString() string {
	return this.DBUserName + ":" + this.DBPassword + "@localhost/" + this.DBName
}

func (this *TTempDB) OpenDB() {
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

func (this *TTempDB) CloseDB() {
	if this.DB != nil {
		this.DB.Close()
	}
}

func (this *TTempDB) PrepareTables() {
	var transaction, _ = this.DB.Begin()
	if CheckTableExists(transaction, "Temperatures") {
	} else {
		transaction.Exec("create table Temperatures (Moment timestamp, temperature float)")
	}
	transaction.Commit()
}