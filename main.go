package main

import (
	"net/http"
	"example/cca-octoprint-backend/httpfunctions"
)

var printers []string = []string{"a", "b", "c", "d", "e", "f"}

func main() {
	http.HandleFunc("/", httpfunctions.MainPage)
	http.HandleFunc("/printer/{printerNumber}/", httpfunctions.PrinterPage)

	http.ListenAndServe(":8080", nil)
}
