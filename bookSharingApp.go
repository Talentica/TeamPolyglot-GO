package main

import (
	"controllers"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	templateFolder, err := os.Open(basePath)
	fmt.Println(err)
	if err != nil {
		return nil
	}
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths,
				basePath+"/"+pathInfo.Name())
			fmt.Println(basePath + "/" + pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)

	return result
}

func Register(templates *template.Template) {
	hc := new(controllers.HomeController)
	hc.Template = templates.Lookup("home.html")
	http.Handle("/", hc)

	lc := new(controllers.LoginController)
	http.Handle("/Login", lc)
}

func main() {

	templates := populateTemplates()
	if templates != nil {
		Register(templates)
		http.ListenAndServe(":8080", nil)
	}

}
