package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/mcuadros/go-octoprint"
)

func main() {
	var printerState octoprint.PrinterState

	printerStateFile, err := os.ReadFile("printer.json")
	if err != nil {
		log.Fatalln("Failed to open printer state file:", err)
	}

	// client := octoprint.NewClient("[BASE PRINTER URL GOES HERE]", "[API KEY GOES HERE]")

	if err := json.Unmarshal(printerStateFile, &printerState); err != nil {
		log.Fatalln("Failed to get printer state:", err)
	}

	fmt.Println(printerState)
}
