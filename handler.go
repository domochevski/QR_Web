package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/skip2/go-qrcode"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles(filepath.Join("templates", "index.html")))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.Execute(w, nil); err != nil {
		log.Printf("template error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	if text == "" {
		http.Error(w, "text parameter is required", http.StatusBadRequest)
		return
	}

	png, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
		log.Printf("qr generation error: %v", err)
		http.Error(w, "Failed to generate QR code", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(png)
}
