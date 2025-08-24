package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/wmarchesi123/go-3dprint-client/octoprint"
)

func getPrinterState(apiKey string) octoprint.PrinterResponse {
	var p octoprint.PrinterResponse

	printerStateFile, err := os.ReadFile("printer.json")
	if err != nil {
		log.Fatalln("Failed to open printer state file:", err)
	}

	if err := json.Unmarshal(printerStateFile, &p); err != nil {
		log.Fatalln("Failed to get printer state:", err)
	}

	return p

}

// A representation of a printer temprerature data entry converted for use in a slice / array. the Name field will contain what was the key for the original temperature struct.
type ModifiedTemperatureData struct {
	Name   string
	Actual float64
	Target float64
	Offset float64
}

type ModifiedPrinterResponse struct {
	PrinterName        string
	State              octoprint.PrinterState
	Temperature        []ModifiedTemperatureData
	TemperatureHistory []octoprint.TemperatureHistory
}

func convertTemperatureData(data octoprint.PrinterResponse, printerName string) ModifiedPrinterResponse {
	var m ModifiedPrinterResponse

	m.PrinterName = fmt.Sprintf("Printer %v", printerName)
	m.State = data.State
	m.TemperatureHistory = data.Temperature.History

	m.Temperature = make([]ModifiedTemperatureData, 6)

	m.Temperature[0].Name = "bed"
	m.Temperature[0].Actual = data.Temperature.Bed.Actual
	m.Temperature[0].Target = data.Temperature.Bed.Target
	m.Temperature[0].Offset = data.Temperature.Bed.Offset

	m.Temperature[1].Name = "tool 0"
	m.Temperature[1].Actual = data.Temperature.Tool0.Actual
	m.Temperature[1].Target = data.Temperature.Tool0.Target
	m.Temperature[1].Offset = data.Temperature.Tool0.Offset

	m.Temperature[2].Name = "tool 1"
	m.Temperature[2].Actual = data.Temperature.Tool1.Actual
	m.Temperature[2].Target = data.Temperature.Tool1.Target
	m.Temperature[2].Offset = data.Temperature.Tool1.Offset

	m.Temperature[3].Name = "tool 2"
	m.Temperature[3].Actual = data.Temperature.Tool2.Actual
	m.Temperature[3].Target = data.Temperature.Tool2.Target
	m.Temperature[3].Offset = data.Temperature.Tool2.Offset

	m.Temperature[4].Name = "tool 3"
	m.Temperature[4].Actual = data.Temperature.Tool3.Actual
	m.Temperature[4].Target = data.Temperature.Tool3.Target
	m.Temperature[4].Offset = data.Temperature.Tool3.Offset

	m.Temperature[5].Name = "tool 4"
	m.Temperature[5].Actual = data.Temperature.Tool4.Actual
	m.Temperature[5].Target = data.Temperature.Tool4.Target
	m.Temperature[5].Offset = data.Temperature.Tool4.Offset

	return m
}
