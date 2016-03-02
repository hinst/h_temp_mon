package h_temp_mon

import (
	"database/sql"
	"strconv"
	"strings"
	"time"
)

func escapeFDBString(text string) string {
	return strings.Replace(text, "'", "''", -1)
}

func CheckTableExists(transaction *sql.Tx, tableName string) bool {
	var result = false
	tableName = escapeFDBString(tableName)
	var queryText = "select 1 from rdb$relations where rdb$relation_name = '" + strings.ToUpper(tableName) + "'"
	var rows, queryError = transaction.Query(queryText)
	if queryError == nil {
		defer rows.Close()
		for rows.Next() {
			result = true
		}
	}
	return result
}

func TimeToFirebirdString(value time.Time) string {
	var a0 = func(value, countOfDigits int) string {
		var text = strconv.Itoa(value)
		for len(text) < countOfDigits {
			text = "0" + text
		}
		return text
	}
	return a0(value.Year(), 4) + "-" + a0(int(value.Month()), 2) + "-" + a0(value.Day(), 2) +
		" " + a0(value.Hour(), 2) + ":" + a0(value.Minute(), 2) + ":" + a0(value.Second(), 2) +
		"." + a0(value.Nanosecond()/100000, 4)
}
