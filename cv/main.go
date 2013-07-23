package cv

import (
    "net/http"
	"html/template"
)

func init() {
    http.HandleFunc("/", root)
	http.HandleFunc("/EnglishCV", EnglishCV)
	http.HandleFunc("/ChineseCV", ChineseCV)
}

func EnglishCV(w http.ResponseWriter, r *http.Request) {
   t, _ := template.ParseFiles("html/EnglishCV.html")
   t.Execute(w, nil)
}

func ChineseCV(w http.ResponseWriter, r *http.Request) {
   t, _ := template.ParseFiles("html/ChineseCV.html")
   t.Execute(w, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/EnglishCV", http.StatusFound)
}

