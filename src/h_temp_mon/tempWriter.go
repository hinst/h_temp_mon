package h_temp_mon

import (
	"database/sql"

	_ "github.com/nakagami/firebirdsql"
)

type TTempWriter struct {
	DBUserName string
	DBPassword string
	DBName     string
	DB         *sql.DB
}

func CreateTempWriter() *TTempWriter {
	var this = &TTempWriter{}
	this.DBUserName = "h_temp_mon"
	this.DBPassword = "password"
	this.DBName = "h_temp_mon"
	return this
}

func (this *TTempWriter) Run() {
	var connectionString = this.GetConnectionString()
	var dbOpenResult error
	this.DB, dbOpenResult = sql.Open("firebirdsql", connectionString)
	if dbOpenResult == nil {

	}
}

func (this *TTempWriter) GetConnectionString() string {
	return this.DBUserName + ":" + this.DBPassword + "@localhost/" + this.DBName
}
