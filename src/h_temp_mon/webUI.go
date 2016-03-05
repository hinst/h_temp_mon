package h_temp_mon

import (
	"net/http"
	"sync"
)

type TWebUI struct {
	Waiter    sync.WaitGroup
	Directory string
	URL       string
}

func CreateWebUI() *TWebUI {
	var result = &TWebUI{}
	return result
}

func (this *TWebUI) Prepare() {
	this.installFileHandler("js")
	this.installFileHandler("css")
}

func (this *TWebUI) installFileHandler(subDirectory string) {
	var url = this.URL + "/" + subDirectory
	var directory = this.Directory + "/" + subDirectory
	Log.Println("installFileHandler:", url, "->", directory)
	http.Handle(url, http.FileServer(http.Dir(directory)))
}

func (this *TWebUI) ProcessRequest(response http.ResponseWriter, request *http.Request) {
}
