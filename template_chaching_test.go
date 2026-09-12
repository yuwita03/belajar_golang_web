package golangweb

import (
	"embed"
	"text/template"
	"net/http"
	"net/http/httptest"
	"testing"
	"io"
	"fmt"
)
//go:embed resources/templates/*.gohtml
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "resources/templates/*.gohtml"))

func TemplateChaching(writer http.ResponseWriter,request *http.Request){
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello Template Chaching")
}

func TestTemplateCaching(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateChaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}