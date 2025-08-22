package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/wmarchesi123/go-3dprint-client/octoprint"
)

func printerStatePage(w http.ResponseWriter, r *http.Response) {
	var printerState octoprint.PrinterResponse

	printerStateFile, err := os.ReadFile("printer.json")
	if err != nil {
		log.Fatalln("Failed to open printer state file:", err)
	}

	if err := json.Unmarshal(printerStateFile, &printerState); err != nil {
		log.Fatalln("Failed to get printer state:", err)
	}

	
}
