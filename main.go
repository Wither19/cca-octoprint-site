package main

import (
	"embed"
	"io/fs"
	"jv/cca-octoprint/httpfunctions"
	"log"
	"net/http"
)

//go:embed frontend/dist/* frontend/dist/assets/*
var staticWebFiles embed.FS

func main() {

	webAppContents, err := fs.Sub(staticWebFiles, "dist")
	if err != nil {
		log.Fatalln("Could not embed static files:", err)
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(webAppContents))))

	http.HandleFunc("/", httpfunctions.MainPage)
	http.HandleFunc("/printer/{printerNumber}/", httpfunctions.PrinterStatePage)
	http.HandleFunc("/api/thing", httpfunctions.APITest)

	http.ListenAndServe(":8080", nil)
}
