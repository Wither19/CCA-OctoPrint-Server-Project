package main

import "net/http"

func main() {
	http.HandleFunc("/printer/{printerNumber}/", printerStatePage)

	http.ListenAndServe(":8080", nil)
}
