package main

import (
	"net/http"
	"fmt"
	"html/template"
	"path/filepath"
)

func main() {
	var config config
	config.fillFromYML("conf.yml")

	var foldersURL []string

	for _, val := range config.Folders {
		folderName := filepath.Base(val)
		foldersURL = append(foldersURL, folderName)

		fs := http.FileServer(http.Dir(val))
		http.Handle("/" + folderName + "/", http.StripPrefix("/" + folderName, fs))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, foldersURL)
	})

	fmt.Println(http.ListenAndServe(config.Port, nil))
}