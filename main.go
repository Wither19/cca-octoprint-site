package main

import (
	"jv/cca-octoprint/httpfunctions"
	"net/http"
)

func main() {
	http.HandleFunc("/", httpfunctions.MainPage)
	http.HandleFunc("/printer/{printerNumber}/", httpfunctions.PrinterStatePage)
	http.HandleFunc("/api/thing", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})

	http.ListenAndServe(":8080", nil)
}
