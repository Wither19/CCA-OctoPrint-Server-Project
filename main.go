package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/wmarchesi123/go-3dprint-client/octoprint"
)

func main() {
	var printerState octoprint.PrinterResponse

	printerStateFile, err := os.ReadFile("printer.json")
	if err != nil {
		log.Fatalln("Failed to open printer state file:", err)
	}

	if err := json.Unmarshal(printerStateFile, &printerState); err != nil {
		log.Fatalln("Failed to get printer state:", err)
	}

	fmt.Println(printerState.State.Text)

}
