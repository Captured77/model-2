package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func params(res http.ResponseWriter, req *http.Request) {
	fmt.Println("entering params")
	version := os.Getenv("VERSION")

	for k, v := range req.Header {
		res.Header().Set(fmt.Sprintf("%v", k), fmt.Sprintf("%v\n", v))
	}
	res.Header().Set("VERSION", version)

}

func healthz(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "OK\n")
}

func main() {
	http.HandleFunc("/", params)
	err := http.ListenAndServe(":8000", nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthz)
	if err != nil {
		log.Fatal(err)
	}
}
