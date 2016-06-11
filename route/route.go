package route

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"

	"github.com/mgerb42/lsstuios/controller"
	"github.com/mgerb42/lsstuios/controller/api"
)

func Routes() *httprouter.Router {

	log.Println("Server Started")

	r := httprouter.New()

	r.GET("/api/:info", api.HandleApiRequest)
	r.GET("/redirect", controller.HandleRedirect)

	//set up public folder path
	r.ServeFiles("/public/*filepath", http.Dir("./public"))

	//route every invalid request to template file
	//routing is all handled on the client side with angular
	r.NotFound = http.HandlerFunc(fileHandler("./public/view/template.html"))

	return r
}

//route requests to static files with httprouter framework
func routerFileHandler(path string) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

		http.ServeFile(w, r, path)

	}
}

//function to serve files with standard net/http library
func fileHandler(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		http.ServeFile(w, r, path)

	}
}
