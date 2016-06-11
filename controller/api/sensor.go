package api

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// handle a simple api request
func HandleApiRequest(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	info := ps.ByName("info")

	//send response back
	fmt.Fprint(w, "{ \"Requested Info\" : \""+info+"\"}")
}
