package main

import (
	"net/http"

	"github.com/Eodomius/CDN/routes"
	"github.com/Eodomius/router"
)


func main(){
	var router = router.New()
	http.Handle("/", router)
	routes.Avatars(&router)
	http.ListenAndServe(":8080", nil)
}