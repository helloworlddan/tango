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
		service.Infof(r, "requested '%s'", r.URL)

		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			service.Criticalf(r, "cannot load template: %v", err)
		}
		tmpl.Execute(w, nil)
	})

	service.ShutdownFunc(func(ctx context.Context, s *run.Service) {
		// TODO: Clean up
	})

	err := service.ListenAndServe()
	if err != nil {
		service.Fatal(nil, err)
	}
}
