package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			w.Write([]byte(fmt.Sprintf("%s = %s\n", k, v)))
		}
		w.Write([]byte(fmt.Sprintf("RemoteAddr = %s\n", r.RemoteAddr)))
		w.WriteHeader(http.StatusOK)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}