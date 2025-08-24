package main

import (
	"fmt"
	"net/http"
	"slices"
	"strings"
	"text/template"
)

func parseTemplate(htmlName string) *template.Template {
	t := template.Must(template.ParseFiles(fmt.Sprintf("./static/pages/%v", htmlName)))

	return t
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	parseTemplate("index.html").Execute(w, printers)
}

func printerStatePage(w http.ResponseWriter, r *http.Request) {
	printerName := r.PathValue("printerNumber")

	if !slices.Contains(printers, strings.ToLower(printerName)) {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		printerState := getPrinterState("[ADD API KEY HERE]")
		printerData := convertTemperatureData(printerState, printerName)

		parseTemplate("printer-overview.html").Execute(w, printerData)

	}

}
