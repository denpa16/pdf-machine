package main

import (
    "fmt"
    "log"
    "net/http"

    wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func urlToPDFHandler(w http.ResponseWriter, r *http.Request) {
    url := r.URL.Query().Get("url")
    format := r.URL.Query().Get("format")
    if url == "" {
        http.Error(w, "Missing 'url' query parameter", http.StatusBadRequest)
        return
    }
    if format == "" {
        http.Error(w, "Missing 'format' query parameter", http.StatusBadRequest)
        return
    }
    pdfg, err := wkhtmltopdf.NewPDFGenerator()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    pdfg.PageSize.Set(format)
    page := wkhtmltopdf.NewPage(url)
    pdfg.AddPage(page)
    err = pdfg.Create()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/pdf")
    _, err = w.Write(pdfg.Bytes())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func main() {
    http.HandleFunc("/convert", urlToPDFHandler)
    port := ":8000"
    fmt.Printf("Listening on port %s...\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatal(err)
    }
}