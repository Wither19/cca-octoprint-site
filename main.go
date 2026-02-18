package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"github.com/wmarchesi123/go-3dprint-client/octoprint"
)

// Initializing two embedded filesystems: one for the HTML and the other for styles and scripts.

//go:embed frontend/dist/*
var staticWebPages embed.FS

//go:embed frontend/dist/assets/*
var staticWebAssets embed.FS

func main() {
	webAppPages, err := fs.Sub(staticWebPages, "dist")
	if err != nil {
		log.Fatalln("Could not embed static files:", err)
	}

	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.FS(webAppPages))))

	webAppAssets, err := fs.Sub(staticWebAssets, "frontend/dist/assets")
	if err != nil {
		log.Fatalln("Could not embed static files:", err)
	}

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.FS(webAppAssets))))

	http.HandleFunc("/", mainPage)
	// http.HandleFunc("/printer/{printerNumber}/", printerStatePage)
	http.HandleFunc("/api/thing", APITest)

	http.ListenAndServe(":8080", nil)
}

func parseTemp(htmlName string, f template.FuncMap) *template.Template {
	templateName := strings.ReplaceAll(htmlName, ".html", "")

	t := template.New(templateName)

	if f != nil {
		t = t.Funcs(f)
	}
	t = template.Must(template.ParseFS(staticWebPages, fmt.Sprintf("frontend/dist/%v", htmlName)))

	return t
}

// var printers []string = []string{"a", "b", "c", "d", "e", "f"}

func mainPage(w http.ResponseWriter, r *http.Request) {
	parseTemp("index.html", nil).Execute(w, nil)
}

// func printerStatePage(w http.ResponseWriter, r *http.Request) {
// 	printerName := r.PathValue("printerNumber")

// 	if !slices.Contains(printers, strings.ToLower(printerName)) {

// 		http.Redirect(w, r, "/", http.StatusMovedPermanently)
// 	} else {
// 		parseTemp("dist/printer-overview.html", nil).Execute(w, nil)
// 	}
// }

func APITest(w http.ResponseWriter, r *http.Request, URL string, APIKey string) {

	p := octoprint.NewClient(APIKey + "api/files/local", APIKey)

	state, err := p.GetPrinterState()
	if err != nil {
		log.Fatalln("printer state: ", err)
	}

	stateJSON, err := json.Marshal(&state)
	if err != nil {
		log.Fatalln("json: ", err)
	}

	w.Write(stateJSON)
}
