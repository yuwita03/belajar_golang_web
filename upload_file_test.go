package golangweb

import (
	"io"
	"net/http"
	"os"
	"testing"
)

func UploadForm(writer http.ResponseWriter,request *http.Request){
	myTemplates.ExecuteTemplate(writer, "upload.form.gohtml",nil)
}

func Upload(writer http.ResponseWriter,request *http.Request){
	// request.ParseMultipartForm(100 << 20)
	file, fileheader, err := request.FormFile("file")
	if err != err {
		panic(err)
	}
	fileDestination, err :=os.Create("./resources/" + fileheader.Filename)
	if err != err {
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	{
		panic(err)
	}
	name := request.PostFormValue("name")
	myTemplates.ExecuteTemplate(writer, "aplod-succed.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileheader.Filename,
	})
}

func TestUploadForm(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/form", UploadForm)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}