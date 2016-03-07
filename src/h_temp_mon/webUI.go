package h_temp_mon

import (
	"bufio"
	"bytes"
	"html/template"
	"io/ioutil"
	"net/http"
	"sync"
)

type TWebUI struct {
	Waiter           sync.WaitGroup
	Directory        string
	URL              string
	PageSubdirectory string
}

type TPageData struct {
	Title  string
	AppURL string
	Body   string
}

func CreateWebUI() *TWebUI {
	var this = &TWebUI{}
	this.PageSubdirectory = "page"
	return this
}

func (this *TWebUI) Prepare() {
	this.installFileHandler("js")
	this.installFileHandler("css")
	this.installTestHandler()
	http.HandleFunc(this.URL+"/status", this.ProcessStatusRequest)
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

func (this *TWebUI) ProcessStatusRequest(response http.ResponseWriter, request *http.Request) {
	var pageData TPageData
	pageData.Title = "Status"
	pageData.AppURL = this.URL
	pageData.Body = this.GetPageContent("status.html")
	response.Write([]byte(this.ApplyTemplate(pageData)))
}

func (this *TWebUI) ApplyTemplate(pageData TPageData) string {
	var preparedTemplate, templateError = template.ParseFiles(this.Directory + "/" + this.PageSubdirectory + "/template.html")
	var result = ""
	if templateError == nil {
		var data bytes.Buffer
		preparedTemplate.Execute(bufio.NewWriter(&data), pageData)
		result = data.String()
	}
	return result
}

func (this *TWebUI) GetPageContent(fileName string) string {
	var data, readResult = ioutil.ReadFile(this.Directory + "/" + this.PageSubdirectory + "/" + fileName)
	var result = ""
	if readResult == nil {
		result = string(data)
	}
	return result
}
