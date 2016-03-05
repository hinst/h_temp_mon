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
	http.Handle(this.URL+"/js", http.FileServer(http.Dir(this.Directory+"/js")))
	http.Handle(this.URL+"/css", http.FileServer(http.Dir(this.Directory+"/css")))
}

func (this *TWebUI) ProcessRequest(response http.ResponseWriter, request *http.Request) {
}
