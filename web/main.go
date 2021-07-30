package main

import (
    "log"
    "net/http"
)


func main()  {
    mux :=http.NewServeMux()
    mux.HandleFunc("/short", longurl)
    mux.HandleFunc("/long", shoturl)
    log.Println("Запуск веб-сервера на http://127.0.0.1:8080")
    err := http.ListenAndServe(":8080", mux)
    log.Fatal(err)
}