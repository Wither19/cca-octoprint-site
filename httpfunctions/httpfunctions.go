package httpfunctions

import (
	"fmt"
	"jv/cca-octoprint/apifunctions"
	"net/http"
	"slices"
	"strings"
	"text/template"
)

var printers []string = []string{"a", "b", "c", "d", "e", "f"}

func ParseTemplate(htmlName string) *template.Template {
	t := template.Must(template.ParseFiles(fmt.Sprintf("./static/pages/%v", htmlName)))

	return t
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	ParseTemplate("index.html").Execute(w, printers)
}

func PrinterStatePage(w http.ResponseWriter, r *http.Request) {
	printerName := r.PathValue("printerNumber")

	if !slices.Contains(printers, strings.ToLower(printerName)) {

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {

		printerState := apifunctions.GetPrinterState("[ADD API KEY HERE]")
		printerData := apifunctions.ConvertTemperatureData(printerState, printerName)

		ParseTemplate("printer-overview.html").Execute(w, printerData)
	}
}
