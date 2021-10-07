package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/healthz", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(200)
	})

	if err := http.ListenAndServe(":28080", log(http.DefaultServeMux)); err != nil {
		fmt.Println("start http server failed.", err)
	}
}

func rootHandler(writer http.ResponseWriter, _ *http.Request) {
	writer.WriteHeader(404)
	fmt.Fprintf(writer, "You Got A Http Server.")
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	name := "World!"
	if params.Has("name") {
		name = params.Get("name")
	}
	resH := writer.Header()
	// test
	resH["foo"] = []string{"bar"}
	// add request header to respose header
	for k, v := range request.Header {
		resH[k] = v
	}

	// print request header to response
	fmt.Fprintf(writer, "Hello, %s!\n", name)
	for k, v := range request.Header {
		fmt.Fprintf(writer, "%s: %s\n", k, v)
	}

}

func log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
