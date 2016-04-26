package main

import (
	"net/http"

	"github.com/Talentica/TeamPolyglot-GO/routers"
)

func main() {
	router := routers.InitRoutes()
	http.ListenAndServe(":5000", router)
}
