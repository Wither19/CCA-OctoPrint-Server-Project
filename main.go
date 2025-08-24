package main

import "net/http"

var printers []string = []string{"a", "b", "c", "d", "e", "f"}

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/printer/{printerNumber}/", printerStatePage)

	http.ListenAndServe(":8080", nil)
}
