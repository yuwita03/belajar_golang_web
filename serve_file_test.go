package golangweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter,request *http.Request) {
	if request.URL.Query().Get("name") != ""{
		http.ServeFile(writer, request, "./resources/ok.html")
	}else{
		http.ServeFile(writer, request, "./resources/err.html")
	}
}

func TestServerFileServer(t *testing.T){
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/ok.html
var resourceOK string
//go:embed resources/err.html
var resourceNotfound string

func ServeFileEmbed(writer http.ResponseWriter,request *http.Request) {
	if request.URL.Query().Get("name") != ""{
		fmt.Fprint(writer, resourceOK)
	}else{
		fmt.Fprint(writer, resourceNotfound)
	}
}

func TestServerFileServerEmbed(t *testing.T){
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
