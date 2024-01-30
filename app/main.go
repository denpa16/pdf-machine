package main

import (
    "fmt"
    "log"
    "net/http"

    wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func urlToPDFHandler(w http.ResponseWriter, r *http.Request) {
    // Get URL from query parameter
    url := r.URL.Query().Get("url")
    if url == "" {
        http.Error(w, "Missing 'url' query parameter", http.StatusBadRequest)
        return
    }

    // Set up the PDF generator
    pdfg, err := wkhtmltopdf.NewPDFGenerator()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Add the URL to convert
    page := wkhtmltopdf.NewPage(url)
    pdfg.AddPage(page)

    // Create and write the PDF to response
    err = pdfg.Create()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/pdf")
    //w.Header().Set("Content-Disposition", "attachment; filename=\"converted.pdf\"")
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