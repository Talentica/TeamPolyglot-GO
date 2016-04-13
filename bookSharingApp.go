package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	basePath := "templates"
	templateFolder, err1 := os.Open(basePath)
	fmt.Println(err1)
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

	result := template.New("templates")
	result, e := result.ParseFiles(*templatePaths...)
	fmt.Println(e) // merge.
	err := result.Lookup("home.html").Execute(w, nil)
	fmt.Println(err) // merge.
}

func main() {

	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)

}
