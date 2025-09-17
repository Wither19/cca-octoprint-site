package apifunctions

import (
	"fmt"
	"log"
	"os"

	"github.com/wmarchesi123/go-3dprint-client/octoprint"
)

func GetPrinterState(APIKey string, baseURL string) []byte {
	printerStateFile, err := os.ReadFile(baseURL)
	if err != nil {
		log.Fatalln("Failed to open printer state file:", err)
	}

	// marshaledPrinterState, err := json.Marshal(&printerStateFile)
	// if err != nil {
	// 	log.Fatalln("Failed to marshal printer state file:", err)
	// }

	return printerStateFile
	// return marshaledPrinterState
}

// func GetPrinterState(apiKey string, baseUrl string) octoprint.PrinterResponse {
// 	var p octoprint.PrinterResponse

// 	printerStateFile, err := os.ReadFile("printer.json")
// 	if err != nil {
// 		log.Fatalln("Failed to open printer state file:", err)
// 	}

// 	if err := json.Unmarshal(printerStateFile, &p); err != nil {
// 		log.Fatalln("Failed to get printer state:", err)
// 	}

// 	return p

// }

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
	StateColor         string
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
	m.StateColor = PrinterStateColors(data.State.Flags)
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

type PrinterStateFlags struct {
	Operational   bool `json:"operational"`
	Paused        bool `json:"paused"`
	Printing      bool `json:"printing"`
	Cancelling    bool `json:"cancelling"`
	Pausing       bool `json:"pausing"`
	SDReady       bool `json:"sdReady"`
	Error         bool `json:"error"`
	Ready         bool `json:"ready"`
	ClosedOrError bool `json:"closedOrError"`
}

func PrinterStateColors(s PrinterStateFlags) string {
	var color string

	if s.Ready && s.SDReady {
		color = "success"
	} else if s.Paused || s.Printing || s.Cancelling || s.Pausing {
		color = "warning"
	} else if s.Error || s.ClosedOrError {
		color = "danger"
	} else {
		color = "secondary"
	}

	return color
}
