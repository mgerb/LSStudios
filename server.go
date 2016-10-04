package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
)

func main() {
    router := httprouter.New()
    router.ServeFiles("/*filepath", http.Dir("./public/"))
    
    log.Println(http.ListenAndServe(":8080", router))
}