package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func printerStatePage(w http.ResponseWriter, r *http.Request) {

	printerState := getPrinterState("[ADD API KEY HERE]")
	printerData := convertTemperatureData(printerState)
	printerData.PrinterName = fmt.Sprintf("Printer %v", r.PathValue("printerNumber"))

	template.Must(template.ParseFiles("./static/pages/printer-overview.html")).Execute(w, printerData)

}
