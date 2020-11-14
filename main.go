package httppages

import (
	"html/template"
	"net/http"
)

// Write404 ...
func Write404(w http.ResponseWriter) {
	t, err := template.New("404").Parse(page404)
	if err != nil {
		w.Write([]byte("Oops, something went seriously wrong!"))
	}
	err = t.ExecuteTemplate(w, "404", nil)
	if err != nil {
		w.Write([]byte("Oops, something went seriously wrong"))
	}
}

// Write401 ...
func Write401(w http.ResponseWriter) {
	t, err := template.New("401").Parse(page401)
	if err != nil {
		w.Write([]byte("Oops, something went seriously wrong!"))
	}
	err = t.ExecuteTemplate(w, "401", nil)
	if err != nil {
		w.Write([]byte("Oops, something went seriously wrong"))
	}
}

// Write415 ...
func Write415(w http.ResponseWriter) {
	t, err := template.New("415").Parse(page415)
	if err != nil {
		w.Write([]byte("Oops, something went seriously wrong!"))
	}
	err = t.ExecuteTemplate(w, "415", nil)
	if err != nil {
		w.Write([]byte("Oops, something went seriously wrong"))
	}
}
