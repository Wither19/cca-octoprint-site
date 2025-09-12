package httpfunctions

import (
	"embed"
	"encoding/json"
	"jv/cca-octoprint/apifunctions"
	"net/http"
	"slices"
	"strings"
	"text/template"
)

var printers []string = []string{"a", "b", "c", "d", "e", "f"}

//go:embed frontend/dist/* frontend/dist/assets/*
var staticWebFiles embed.FS

func ParseTemplate(htmlName string, f template.FuncMap) *template.Template {
	templateName := strings.ReplaceAll(htmlName, ".html", "")

	t := template.New(templateName)

	if f != nil {
		t = t.Funcs(f)
	}

	t = template.Must(template.ParseFS(main.staticWebFiles, htmlName))

	return t
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	ParseTemplate("/static/index.html", nil).Execute(w, nil)
}

func PrinterStatePage(w http.ResponseWriter, r *http.Request) {
	printerName := r.PathValue("printerNumber")

	if !slices.Contains(printers, strings.ToLower(printerName)) {

		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	} else {
		ParseTemplate("static/printer-overview.html", nil).Execute(w, nil)
	}
}

func APITest(w http.ResponseWriter, r *http.Request) {

	printerResponse := apifunctions.ConvertTemperatureData(apifunctions.GetPrinterState("", ""), "")

	APIMsg, _ := json.Marshal(printerResponse)

	w.Write(APIMsg)

}
