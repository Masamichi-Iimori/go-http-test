package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi, I'm masamichhhhhhhhhhhhhhhhhhi!\n")
	})
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "you reached /test path\n")
	})
	http.ListenAndServe(":80", nil)
}
