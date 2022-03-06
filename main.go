package main

import (
	"net/http"

	"github.com/Eodomius/CDN/routes"
	"github.com/Eodomius/router"
)

const PORT = "8080"

func main(){

	var router = router.New()
	http.Handle("/", router)
	routes.Avatars(&router)

	http.ListenAndServe(":"+PORT, nil)
}