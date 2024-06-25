package main

import (
	"context"
	"net/http"
	"text/template"

	"github.com/helloworlddan/run"
)

func main() {
	http.Handle(
		"/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))),
	)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		run.Infof(r, "requested '%s'", r.URL)

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			run.Criticalf(r, "cannot load template: %v", err)
		}
		tmpl.Execute(w, nil)
	})

	shutdown := func(ctx context.Context) {
		// TODO: Clean up
	}

	err := run.ServeHTTP(shutdown, nil)
	if err != nil {
		run.Fatal(nil, err)
	}
}
