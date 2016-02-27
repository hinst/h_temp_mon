package main

import (
	"fmt"
	"h_temp_mon"
)

func main() {
	fmt.Println("[GLOBAL START]")
	var app = &h_temp_mon.TApp{}
	app.Run()
}
