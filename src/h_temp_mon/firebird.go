package h_temp_mon

import "database/sql"
import "strings"

func escapeFDBString(text string) string {
	return strings.Replace(text, "'", "''", -1)
}

func CheckTableExists(transaction *sql.Tx, tableName string) bool {
	tableName = escapeFDBString(tableName)
	var x = 1
	var row = transaction.QueryRow("select 1 from rdb$relations where rdb$relation_name = '" + tableName + "'")
	var scanResult = row.Scan(&x)
	return scanResult == nil
}
