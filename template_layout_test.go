package go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml", "./templates/footer.gohtml", "./templates/layout.gohtml",
	))
	t.ExecuteTemplate(w, "layout", map[string]string{
		"Title": "Template Data Map",
		"Name":  "Akbar",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
