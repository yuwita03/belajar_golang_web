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

func Upload(writer http.ResponseWriter, request *http.Request) {
    if err := request.ParseMultipartForm(10 << 20); err != nil {
        panic(err)
    }

    file, fileHeader, err := request.FormFile("file")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    dst, err := os.Create("./resources/" + fileHeader.Filename)
    if err != nil {
        panic(err)
    }
    defer dst.Close()

    _, err = io.Copy(dst, file)
    if err != nil {
        panic(err)
    }

    name := request.PostFormValue("name")

    myTemplates.ExecuteTemplate(writer, "aplod-succed.gohtml", map[string]interface{}{
        "Name": name,
        "File": "/static/" + fileHeader.Filename,
    })
}

func TestUploadForm(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/form", UploadForm)
	mux.HandleFunc("/apload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}