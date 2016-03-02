package h_temp_mon

import "database/sql"
import "strings"

func escapeFDBString(text string) string {
	return strings.Replace(text, "'", "''", -1)
}

func CheckTableExists(transaction *sql.Tx, tableName string) bool {
	var result = false
	tableName = escapeFDBString(tableName)
	var queryText = "select 1 from rdb$relations where rdb$relation_name = '" + strings.ToUpper(tableName) + "'"
	Log.Println(queryText)
	var rows, queryError = transaction.Query(queryText)
	if queryError == nil {
		defer rows.Close()
		for rows.Next() {
			result = true
		}
	}
	return result
}
