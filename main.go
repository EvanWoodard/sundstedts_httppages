package errorpage

import (
	"html/template"
	"net/http"
)

// Write404 ...
func Write404(w http.ResponseWriter) {
	t, err := template.ParseFiles("404.gohtml")
	if err != nil {
		w.Write([]byte("Oops, something went seriously wrong!"))
	}
	err = t.ExecuteTemplate(w, "404", nil)
	if err != nil {
		w.Write([]byte("Oops, something went seriously wrong"))
	}
}
