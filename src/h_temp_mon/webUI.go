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

type TPageStruct struct {
	Title  string
	AppURL string
	Body   string
}

func CreateWebUI() *TWebUI {
	var result = &TWebUI{}
	return result
}

func (this *TWebUI) Prepare() {
	this.installFileHandler("js")
	this.installFileHandler("css")
	this.installTestHandler()
}

func (this *TWebUI) installFileHandler(subDirectory string) {
	var url = this.URL + "/" + subDirectory + "/"
	var directory = this.Directory + "/" + subDirectory
	Log.Println("installFileHandler: '" + url + "' -> '" + directory + "'")
	var fileServer = http.FileServer(http.Dir(directory))
	http.Handle(url, http.StripPrefix(url, fileServer))
}

func (this *TWebUI) installTestHandler() {
	var url = this.URL + "/test"
	Log.Println("test handler installed at '" + url + "'")
	http.HandleFunc(url, this.processTestRequest)
}

func (this *TWebUI) processTestRequest(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("test"))
}

func (this *TWebUI) ProcessRequest(response http.ResponseWriter, request *http.Request) {
}
