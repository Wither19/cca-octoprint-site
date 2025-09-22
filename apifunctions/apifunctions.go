package apifunctions

import (
	"log"
	"os"
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
