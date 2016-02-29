package h_temp_mon

import (
	"io"
	"log"
	"os"
)

var Log *log.Logger
var LogWriter io.Writer

func InitializeLog() {
	LogWriter = os.Stdout
	Log = log.New(LogWriter, "h_temp_mon", log.Lshortfile|log.LstdFlags)
}
