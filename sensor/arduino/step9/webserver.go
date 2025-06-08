package main

import (
	"fmt"
	"net/http"
	"time"

	_ "embed"
)

//go:embed index.html
var page []byte

// Uses Min - a tiny framework that makes websites pretty.
// See https://mincss.com/
//
//go:embed mincss.min.css
var mincss []byte

//go:embed tetromino.html
var tetromino []byte

const port = ":8080"

func startWebServer() {
	http.HandleFunc("/", root)
	http.HandleFunc("/mincss.min.css", css)
	http.HandleFunc("/6", sixlines)
	http.HandleFunc("/on", systemActivate)
	http.HandleFunc("/off", systemDeactivate)
	http.HandleFunc("/status", systemStatus)

	err := http.ListenAndServe(port, nil)
	for err != nil {
		fmt.Printf("error: %s\r\n", err.Error())
		time.Sleep(1 * time.Second)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(page)
}

func css(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write(mincss)
}

// https://fukuno.jig.jp/3267
func sixlines(w http.ResponseWriter, r *http.Request) {
	w.Write(tetromino)
}

func systemActivate(w http.ResponseWriter, r *http.Request) {
	systemActive = true
	w.Header().Set(`Content-Type`, `text/plain; charset=UTF-8`)
	fmt.Fprintf(w, "system active")
}

func systemDeactivate(w http.ResponseWriter, r *http.Request) {
	systemActive = false
	w.Header().Set(`Content-Type`, `text/plain; charset=UTF-8`)
	fmt.Fprintf(w, "system inactive")
}

func systemStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Content-Type`, `text/plain; charset=UTF-8`)
	status := "inactive"
	if systemActive {
		status = "active"
	}
	w.Header().Set(`Content-Type`, `application/json`)
	fmt.Fprintf(w, `{"status": "%s"}`, status)
}
