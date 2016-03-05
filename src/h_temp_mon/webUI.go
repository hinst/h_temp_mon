package h_temp_mon

import (
	"net/http"
	"sync"
)

type TWebUI struct {
	Waiter sync.WaitGroup
}

func (this *TWebUI) ProcessRequest(response http.ResponseWriter, request *http.Request) {

}
