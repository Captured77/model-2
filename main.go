package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func params(res http.ResponseWriter, req *http.Request) {
	fmt.Println("entering params")
	version := os.Getenv("VERSION")
	res.Header().Set("VERSION", version)
	for k, v := range req.Header {
		for _, vv := range v {
			res.Header().Set(k, vv)
		}
	}
	reqIP := getIp(req)
	fmt.Printf("Remote IP: %s, Resp Status Code: %d\n", reqIP, http.StatusOK)
	log.Printf("Remote IP: %s\n", reqIP)
	log.Printf("Resp Status Code: %d\n", http.StatusOK)

}

func getIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr

	if ip := req.Header.Get("X-REAL-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

func healthz(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	io.WriteString(res, "OK\n")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/", params)
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
