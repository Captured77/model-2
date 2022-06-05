package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("entering root handler")
	for k, v := range req.Header {
		res.Header().Set(fmt.Sprintf("%v", k), fmt.Sprintf("%v\n", v))
	}

}

func main() {

	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
