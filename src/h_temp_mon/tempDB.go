package h_temp_mon

import "database/sql"

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
