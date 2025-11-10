package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/mesh-dell/unit-converter/internal/converters"
)

func WeightHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]any{
		"Page":   "weight",
		"Result": 0.0,
	}
	if r.Method == http.MethodPost {
		value := r.PostFormValue("value")
		from := r.PostFormValue("from")
		to := r.PostFormValue("to")

		if value != "" {
			result, err := converters.ConvertWeight(value, from, to)
			if err != nil {
				fmt.Println(err)
			}
			data["Result"] = result
		}
	}
	tmpl := template.Must(template.ParseFiles("internal/templates/base.html", "internal/templates/weight.html"))
	tmpl.ExecuteTemplate(w, "base.html", data)
}
