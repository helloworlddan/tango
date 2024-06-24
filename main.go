package main

import (
	"context"
	"net/http"
	"text/template"

	"github.com/helloworlddan/run"
)

func main() {
	service := run.NewService()

	service.HandleStatic("/static/", "./static/")

	service.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		run.Infof(r, "requested '%s'", r.URL)

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			run.Criticalf(r, "cannot load template: %v", err)
		}
		tmpl.Execute(w, nil)
	})

	service.ShutdownFunc(func(ctx context.Context) {
		// TODO: Clean up
	})

	err := service.ListenAndServeHTTP()
	if err != nil {
		run.Fatal(nil, err)
	}
}
