package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc(`/`, indexHandler)
	http.ListenAndServe(":3000", nil)
}

func indexHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Message")
}
