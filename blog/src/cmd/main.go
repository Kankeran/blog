package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/Kankeran/blog-service/blog/src/internal/common"
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
	path := fmt.Sprintf("%s/templates/home.html", common.ProvideWorkdir())
	parsedTemplate, _ := template.ParseFiles(path)
	err := parsedTemplate.Execute(w, ps.ByName("name"))
	check(err)
}

func main() {
	fmt.Println("Blog run!!!")
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/wiktror/:name", Wiktror)
	path := fmt.Sprintf("%s/static", common.ProvideWorkdir())
	router.ServeFiles("/static/*filepath", http.Dir(path))

	log.Fatal(http.ListenAndServe(":8080", router))
}
