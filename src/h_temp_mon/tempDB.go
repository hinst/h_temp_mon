package h_temp_mon

import (
	"database/sql"
	"io/ioutil"
	"strconv"
	"time"
)

type TTempDB struct {
	DBUserName   string
	DBPassword   string
	DBName       string
	DB           *sql.DB
	DBErrorCount int
	DBErrorLimit int
}

func CreateTempDB() *TTempDB {
	this := &TTempDB{}
	this.DBUserName = "h_temp_mon"
	this.DBPassword = "password"
	this.DBName = "h_temp_mon"
	this.DB = nil
	this.DBErrorCount = 0
	this.DBErrorLimit = 100
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

func (this *TTempDB) Open() {
	var connectionString = this.GetConnectionString()
	var dbOpenResult error
	Log.Println("OpenDB: DBName=" + this.DBName)
	this.DB, dbOpenResult = sql.Open("firebirdsql", connectionString)
	if dbOpenResult == nil {
		var pingResult = this.DB.Ping()
		if pingResult == nil {
			Log.Println("OpenDB: success")
		} else {
			Log.Println("OpenDB: ping failed; error='" + pingResult.Error() + "'")
			this.DB = nil
		}
	} else {
		Log.Println("OpenDB: fail: ", dbOpenResult)
		this.DB = nil
	}
}

func (this *TTempDB) Close() {
	if this.DB != nil {
		var closeResult = this.DB.Close()
		if closeResult == nil {
			Log.Println("DB closed successfully")
		} else {
			Log.Println("Could not close DB; error='" + closeResult.Error() + "'")
		}
	}
}

func (this *TTempDB) PrepareTables() {
	var transaction, _ = this.DB.Begin()
	this.PrepareTable(transaction, "Temperatures", "moment timestamp, temperature float")
	transaction.Commit()
}

func (this *TTempDB) PrepareTable(transaction *sql.Tx, tableName, tableFields string) {
	if CheckTableExists(transaction, tableName) {
	} else {
		var _, err = transaction.Exec("create table " + tableName + " (" + tableFields + ")")
		if err == nil {
			Log.Println("Table created: " + tableName)
		} else {
			Log.Println("Could not create table: " + tableName + " error='" + err.Error() + "'")
		}
	}
}

func (this *TTempDB) WriteTemperature(temperature float32) {
	var momentString = "'" + TimeToFirebirdString(time.Now()) + "'"
	var temperatureString = strconv.FormatFloat(float64(temperature), 'f', 1, 32)
	var statementString = "insert into Temperatures values(" + momentString + ", " + temperatureString + ")"
	var _, err = this.DB.Exec(statementString)
	if err == nil {
	} else {
		this.ReportError("WriteTemperature: error='" + err.Error() + "' statement='" + statementString + "'")
	}
}

func (this *TTempDB) ReadLatestMoment(transaction *sql.Tx) time.Time {
	var result time.Time
	var row = transaction.QueryRow("select max(moment) from Temperatures")
	var queryResult = row.Scan(&result)
	if queryResult != nil {
		this.ReportError(queryResult.Error())
	}
	return result
}

func (this *TTempDB) ReadLatestTemperature() TTempDBRow {
	var result TTempDBRow
	var row = this.DB.QueryRow("select * from Temperatures order by moment desc rows 1")
	var rowError = row.Scan(&result.Moment, &result.Temperature)
	if rowError != nil {
		this.ReportError(rowError.Error())
	}
	return result
}

func (this *TTempDB) ReportError(text string) {
	if this.DBErrorCount < this.DBErrorLimit {
		Log.Println("DB Error #" + strconv.Itoa(this.DBErrorCount) + ": " + text)
	} else {
		Log.Panicln("DB Error #" + strconv.Itoa(this.DBErrorCount) + " limit reached: " + text)
	}
	this.DBErrorCount++
}
