package main

import (
    "log"
    "net/http"
    handler "ascii_web/handlers"
)

var port string = ":8000"

func main() {
    router := http.NewServeMux()
    router.HandleFunc("/",handler.Home)
    router.HandleFunc("/ascii-art",handler.Form)
    log.Printf("\x1b[0;32mproject runing in 'http://localhost%s'.\x1b[0m\n",port)
    err := http.ListenAndServe(port,router)
    log.Fatal(err)
}
