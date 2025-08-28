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

func ParseTemplate(htmlName string, f template.FuncMap) *template.Template {
	templateName := strings.ReplaceAll(htmlName, "./static/pages/", "")
	templateName = strings.ReplaceAll(templateName, ".html", "")

	t := template.New(templateName)

	if f != nil {
		t = t.Funcs(f)
	}

	t = template.Must(template.ParseFiles(fmt.Sprintf("./static/pages/%v", htmlName)))

	return t
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	ParseTemplate("index.html", nil).Execute(w, printers)
}

func PrinterStatePage(w http.ResponseWriter, r *http.Request) {
	printerName := r.PathValue("printerNumber")

	if !slices.Contains(printers, strings.ToLower(printerName)) {

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {

		printerState := apifunctions.GetPrinterState("<API KEY>", "<BASE URL>")
		printerData := apifunctions.ConvertTemperatureData(printerState, printerName)

		ParseTemplate("printer-overview.html", template.FuncMap{"printerStateColors": apifunctions.PrinterStateColors}).Execute(w, printerData)
	}
}
