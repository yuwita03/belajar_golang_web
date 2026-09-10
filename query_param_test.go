package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request){
	name := request.URL.Query().Get("name")

	if name == ""{
		fmt.Fprint(writer, "Hello")
	}else{
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestQueryParameter(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Neko", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameter(writer http.ResponseWriter, request *http.Request){
	firstname := request.URL.Query().Get("first_name")
	lastname := request.URL.Query().Get("last_name")
	
	fmt.Fprintf(writer, "Hello %s %s", firstname, lastname)
}

func TestMultipleQuery(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=Neko&last_name=Koyu", nil)
	recorder := httptest.NewRecorder()

	MultipleParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameterValue(writer http.ResponseWriter,request *http.Request){
	query:= request.URL.Query()
	names := query["name"]


	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Neko&name=Koyu&name=rico", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValue(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}