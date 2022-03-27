package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/input", HundlerUserFunc)
	http.HandleFunc("/view", HandlerUserConfirm)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func HundlerUserFunc(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("template/input.html"))

	values := map[string]string{}

	if err := tpl.ExecuteTemplate(w, "input.html", values); err != nil {
		fmt.Println(err)
	}
}

func HandlerUserConfirm(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("template/view.html"))

	values := map[string]string{
		"name":    r.FormValue("name"),
		"comment": r.FormValue("comment"),
	}

	if err := tpl.ExecuteTemplate(w, "view.html", values); err != nil {
		fmt.Println(err)
	}
}
