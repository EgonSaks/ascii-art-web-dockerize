package main

import (
	"net/http"
	"strconv"
)

type TemplateData struct {
	Output string
	Input  string
}

func download(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/download" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		r.ParseForm()
		input := r.FormValue("textarea_input")
		if !inputValidation(input) {
			http.Error(w, "400 - bad request", http.StatusBadRequest)
			return
		}

		fonts := r.FormValue("font")
		var banner string
		switch fonts {
		case "standard":
			banner = "banner/standard.txt"
		case "shadow":
			banner = "banner/shadow.txt"
		case "tinkertoy":
			banner = "banner/tinkertoy.txt"
		default:
			http.Error(w, "Please go back and choose the font!\n", http.StatusInternalServerError)
		}

		font, err := mapFonts(banner)
		if err != nil {
			http.Error(w, "500 - internal server error", http.StatusInternalServerError)
			return
		}

		output, err := mapUserInput(input, font)
		if err != nil {
			http.Error(w, "400 - bad request", http.StatusBadRequest)
			return
		}

		mime := http.DetectContentType([]byte(output))
		fileSize := len(string(output))

		extension := r.FormValue("download")

		if (extension == ".txt") || (extension == ".doc") {
			w.Header().Set("Content-Disposition", "attachment; filename=ascii-art-result"+extension)
			w.Header().Set("Content-Type", mime)
			w.Header().Set("Content-Length", strconv.Itoa(fileSize))
			_, err := w.Write([]byte(output))
			if err != nil {
				http.Error(w, "500 - internal server error", http.StatusInternalServerError)
				return
			}
		}

		data := TemplateData{
			Output: output,
			Input:  input,
		}

		tpl.ExecuteTemplate(w, "index.html", data)
	default:
		http.Error(w, "405 - bad request", http.StatusMethodNotAllowed)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return

	}

	switch r.Method {
	case "GET":
		tpl.ExecuteTemplate(w, "index.html", nil)
	default:
		http.Error(w, "405 - bad request", http.StatusMethodNotAllowed)
	}
}
