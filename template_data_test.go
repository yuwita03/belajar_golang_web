package golangweb


import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type Address struct{
	Country string
}

type Page struct {
	Title string
	Name string
	Address Address
}

func TemplateDataMap(writer http.ResponseWriter,request *http.Request){
	t := template.Must(template.ParseFiles("./resources/templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title : "Template Data Map",
		Name : "Nekoyu",
		Address: Address {
			Country: "Japan",
		},
	})
}

func TestTemplateDataMap(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}


func TemplateActionIf(writer http.ResponseWriter,request *http.Request){
	t := template.Must(template.ParseFiles("./resources/templates/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", Page{
		Title : "Template Data Map",
		// Name : "Nekoyu",
		Address: Address {
			Country: "Japan",
		},
	})
}

func TestTemplateActionIf(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateOperator(writer http.ResponseWriter,request *http.Request){
	t := template.Must(template.ParseFiles("./resources/templates/comparator.gohtml"))
	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"FinalValue" : 90,
	},)
}

func TestTemplateOperator(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateRange(writer http.ResponseWriter,request *http.Request){
	t := template.Must(template.ParseFiles("./resources/templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Hobbies": []string{
			"Game", "Read", "Code",
		},
	},)
}

func TestTemplateRange(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateWith(writer http.ResponseWriter,request *http.Request){
	t := template.Must(template.ParseFiles("./resources/templates/address.gohtml"))
	t.ExecuteTemplate(writer, "address.gohtml", map[string]interface{}{
		"Title" : "Template ",
		"Name"	: "Neko",
		"Address" : map[string]interface{}{
			"Street": "Yahama",
			"City"	:	"Tokyo",
		},
	},)
}

func TestTemplateWith(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateLayout(writer http.ResponseWriter,request *http.Request){
	t := template.Must(template.ParseFiles(
		"./resources/templates/footer.gohtml",
		"./resources/templates/header.gohtml",
		"./resources/templates/layout.gohtml",
		))
	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title" : "Template",
		"Name"	: "Neko",
		"Address" : map[string]interface{}{
			"Street": "Yahama",
			"City"	:	"Tokyo",
		},
	} ,)
}

func TestTemplateLayout(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
