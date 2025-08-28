package main

import (
	"jv/cca-octoprint/httpfunctions"
	"net/http"
)

func main() {
	http.HandleFunc("/", httpfunctions.MainPage)
	http.HandleFunc("/printer/{printerNumber}/", httpfunctions.PrinterStatePage)

	http.ListenAndServe(":8080", nil)
}
