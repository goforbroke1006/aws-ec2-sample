package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(200)
		_, _ = w.Write([]byte("ok"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err.Error())
	}
}
