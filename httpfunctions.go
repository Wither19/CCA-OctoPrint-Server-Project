package main

import (
	"net/http"
	"text/template"
)

func printerStatePage(w http.ResponseWriter, r *http.Request) {

	printerState := convertTemperatureData(getPrinterState("[ADD API KEY HERE]"))

	template.Must(template.ParseFiles("./static/pages/printer-overview.html")).Execute(w, printerState)

}
