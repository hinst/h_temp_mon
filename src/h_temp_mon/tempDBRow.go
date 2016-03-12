package h_temp_mon

import "time"

type TTempDBRow struct {
	Temperature float32   `xml:"Temperature"`
	Moment      time.Time `xml:"Moment"`
}
