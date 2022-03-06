package routes

import (
	"net/http"

	"github.com/Eodomius/router"
)

func Test(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World"))
}

func Avatars(router *router.Router){
	router.Get("/avatars", Test)
}