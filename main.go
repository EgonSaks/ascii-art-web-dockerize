package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*html"))
}

const (
	port = ":8080"
	host = "http://localhost"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/download", download)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting application on port %s, to shut it down press "+"\033[10;31m"+"CTRL + c"+"\033[0m", port)
	log.Println("Open application:", "\033[10;32m"+host+port+"\033[0m")

	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("%v - internal server error", http.StatusInternalServerError)
	}
}
