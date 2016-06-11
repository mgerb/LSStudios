package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Redirect to discord
func HandleRedirect(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	http.Redirect(w, r, "http://google.com", 301)

}
