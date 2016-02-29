package h_temp_mon

import (
	"io"
	"log"
)

var Log *log.Logger
var LogWriter io.Writer

func InitializeLog() {

	Log = log.New(LogWriter, "h_temp_mon", log.Lshortfile|log.LstdFlags)
}
