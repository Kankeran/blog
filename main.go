package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Wiktror(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	parsedTemplate, _ := template.ParseFiles("./home.html")
	err := parsedTemplate.Execute(w, ps.ByName("name"))
	check(err)
}

func main() {
	fmt.Println("Blog run!!!")
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/wiktror/:name", Wiktror)
	router.ServeFiles("/static/*filepath", http.Dir("./static"))

	log.Fatal(http.ListenAndServe(":8080", router))
}
