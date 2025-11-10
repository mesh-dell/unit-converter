package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/mesh-dell/unit-converter/internal/converters"
)

func TemperatureHandler(w http.ResponseWriter, r *http.Request) {

	data := map[string]any{
		"Page":   "temperature",
		"Result": 0.0,
	}
	if r.Method == http.MethodPost {
		value := r.PostFormValue("value")
		from := r.PostFormValue("from")
		to := r.PostFormValue("to")

		if value != "" {
			result, err := converters.ConvertTemp(value, from, to)
			if err != nil {
				fmt.Println(err)
			}
			data["Result"] = result
		}
	}
	tmpl := template.Must(template.ParseFiles("internal/templates/base.html", "internal/templates/temperature.html"))
	tmpl.ExecuteTemplate(w, "base.html", data)
}
