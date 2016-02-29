package h_temp_mon

import (
	"database/sql"

	_ "github.com/nakagami/firebirdsql"
)

type TTempWriter struct {
	DB *sql.DB
}

func (this *TTempWriter) Run() {
	this.DB = sql.Open("firebirdsql")
}
