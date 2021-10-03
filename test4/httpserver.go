package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my Gosite!!\n\n")
	for k,v := range r.Header {
		fmt.Fprintf(w, "%s: %s\n",k,v)
	}
    fmt.Fprintf(w,"\n\n")
	VERSION := os.Getenv("VERSION")
	w.Header().Set("VERSION",VERSION)
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "response'header VERSION: %s\n",w.Header().Get("VERSION"))
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		fmt.Println("client ip is: %s, response status code: %s",ip)
	}
	fmt.Println("server response statuscode is: ",http.StatusOK)
}
func indexHandler_healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w,"Healthz returns %s (200)\n",http.StatusText(200))
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/healthz", indexHandler_healthz)

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal("failed to start server",err)
	}

}