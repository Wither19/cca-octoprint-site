package main

import (
	"embed"
	"jv/cca-octoprint/httpfunctions"
	"net/http"
)

func main() {

	//go:embed ../fbuild/*
	var staticWebFiles embed.FS

	http.HandleFunc("/", httpfunctions.MainPage)
	http.HandleFunc("/printer/{printerNumber}/", httpfunctions.PrinterStatePage)
	http.HandleFunc("/api/thing", httpfunctions.APITest)

	http.ListenAndServe(":8080", nil)
}
