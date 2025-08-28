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
	t := template.Must(template.ParseFiles(fmt.Sprintf("./static/pages/%v", htmlName)))

	if f != nil {
		t = t.Funcs(f)
	}

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

		printerState := apifunctions.GetPrinterState("[ADD API KEY HERE]")
		printerData := apifunctions.ConvertTemperatureData(printerState, printerName)

		printerStateFuncMap := map[string]any{
			"printerStateColors": func(stateFlags apifunctions.PrinterStateFlags) string {
				var color string
				
				if stateFlags.Ready && stateFlags.SDReady {
					color = "success"
				} else if stateFlags.Paused || stateFlags.Printing || stateFlags.Cancelling || stateFlags.Pausing {
					color = "warning"
				} else if stateFlags.Error || stateFlags.ClosedOrError {
					color = "danger"
				} else {
					color = "secondary"
				}

				return color
			},
		}

		ParseTemplate("printer-overview.html", printerStateFuncMap).Execute(w, printerData)
	}
}
