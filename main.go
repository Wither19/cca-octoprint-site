package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"jv/cca-octoprint/apifunctions"
	"log"
	"net/http"
)

//go:embed frontend/dist/*
var staticWebFiles embed.FS

func main() {
	webAppContents, err := fs.Sub(staticWebFiles, "dist")
	if err != nil {
		log.Fatalln("Could not embed static files:", err)
	}

	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.FS(webAppContents))))

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/printer/{printerNumber}/", printerStatePage)
	http.HandleFunc("/api/thing", APITest)

	http.ListenAndServe(":8080", nil)
}

// func parseTemp(htmlName string, f template.FuncMap) *template.Template {
// 	templateName := strings.ReplaceAll(htmlName, ".html", "")

// 	t := template.New(templateName)

// 	if f != nil {
// 		t = t.Funcs(f)
// 	}

// 	t = template.Must(template.ParseFS(staticWebFiles, htmlName))

// 	return t
// }

var printers []string = []string{"a", "b", "c", "d", "e", "f"}

func mainPage(w http.ResponseWriter, r *http.Request) {
	template.parseTemplate("dist/index.html", nil).Execute(w, nil)
}

// func printerStatePage(w http.ResponseWriter, r *http.Request) {
// 	printerName := r.PathValue("printerNumber")

// 	if !slices.Contains(printers, strings.ToLower(printerName)) {

// 		http.Redirect(w, r, "/", http.StatusMovedPermanently)
// 	} else {
// 		parseTemp("dist/printer-overview.html", nil).Execute(w, nil)
// 	}
// }

func APITest(w http.ResponseWriter, r *http.Request) {

	printerResponse := apifunctions.ConvertTemperatureData(apifunctions.GetPrinterState("", ""), "")

	APIMsg, _ := json.Marshal(printerResponse)

	w.Write(APIMsg)

}
