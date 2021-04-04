package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi, I'm masamichhhi!\n")
	})
	http.ListenAndServe(":80", nil)
}
