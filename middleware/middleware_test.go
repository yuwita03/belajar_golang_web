package middleware

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

type ErrorHandler struct{
	Hanndler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter,request *http.Request){
	fmt.Println("Before Middleware")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Middleware")

}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter,request *http.Request){
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("terjadi error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	errorHandler.Hanndler.ServeHTTP(writer, request)
}
func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request){
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})
	mux.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request){
		fmt.Println("Foo Executed")
		fmt.Fprint(writer, "Foo Middleware")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request){
		fmt.Fprint(writer, "panic Middleware")
		panic("ups")
	})
	LogMiddleware := &LogMiddleware{
		Handler: mux,
	}

	ErrorHandler:=&ErrorHandler{
		Hanndler: LogMiddleware,
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: ErrorHandler,
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}
