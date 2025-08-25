package main

import (
	"example/cca-octoprint-backend/httpfunctions"
	"net/http"
)

var printers []string = []string{"a", "b", "c", "d", "e", "f"}

func main() {
	http.HandleFunc("/", httpfunctions.MainPage)
	http.HandleFunc("/printer/{printerNumber}/", httpfunctions.PrinterStatePage)

	http.ListenAndServe(":8080", nil)
}
