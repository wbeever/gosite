package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/services", services)
	http.HandleFunc("/gallery", gallery)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/mobile", mobile)
	http.HandleFunc("/mgall", mgall)
	http.HandleFunc("/mabou", mabou)
	http.HandleFunc("/mcont", mcont)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", nil)
}
func services(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "services.html", nil)
}
func gallery(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "gallery.html", nil)
}
func about(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.html", nil)
}
func contact(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "contact.html", nil)
}
func mobile(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "mobile.html", nil)
}
func mgall(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "mgall.html", nil)
}
func mabou(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "mabou.html", nil)
}
func mcont(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "mcont.html", nil)
}
