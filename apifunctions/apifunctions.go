package apifunctions

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/wmarchesi123/go-3dprint-client/octoprint"
)

func GetPrinterState(apiKey string) octoprint.PrinterResponse {
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

func fillInTemperatureData(name string, original octoprint.TemperatureData) ModifiedTemperatureData {
	var new ModifiedTemperatureData

	new.Name = name
	new.Actual = original.Actual
	new.Target = original.Target
	new.Offset = original.Offset

	return new
}

func ConvertTemperatureData(data octoprint.PrinterResponse, printerName string) ModifiedPrinterResponse {
	var m ModifiedPrinterResponse

	m.PrinterName = fmt.Sprintf("Printer %v", printerName)
	m.State = data.State
	m.TemperatureHistory = data.Temperature.History

	m.Temperature = make([]ModifiedTemperatureData, 6)

	m.Temperature[0] = fillInTemperatureData("bed", data.Temperature.Bed)
	m.Temperature[1] = fillInTemperatureData("tool-0", data.Temperature.Tool0)
	m.Temperature[2] = fillInTemperatureData("tool-1", data.Temperature.Tool1)
	m.Temperature[3] = fillInTemperatureData("tool-2", data.Temperature.Tool2)
	m.Temperature[4] = fillInTemperatureData("tool-3", data.Temperature.Tool3)
	m.Temperature[5] = fillInTemperatureData("tool-4", data.Temperature.Tool4)

	return m
}
