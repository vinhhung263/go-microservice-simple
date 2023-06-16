package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const portFE = "8081"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "test.page.gohtml")
	})

	log.Printf("Starting front end service on port %s\n", portFE)
	err := http.ListenAndServe(fmt.Sprintf(":%s", portFE), nil)
	if err != nil {
		log.Panic(err)
	}
}

//go:embed templates
var templateFS embed.FS

func render(w http.ResponseWriter, t string) {

	partials := []string{
		"templates/base.layout.gohtml",
		"templates/header.partial.gohtml",
		"templates/footer.partial.gohtml",
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("templates/%s", t))

	for _, x := range partials {
		templateSlice = append(templateSlice, x)
	}

	tmpl, err := template.ParseFS(templateFS, templateSlice...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data struct {
		BrokerURL string
	}

	// data.BrokerURL = os.Getenv("BROKER_URL")
	data.BrokerURL = "http://localhost:8080/"

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
