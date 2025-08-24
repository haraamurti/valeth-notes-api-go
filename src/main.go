package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
)

//handler welcome to localhost:9000
func handlerwelcome (w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}

var data = map[string]interface{}{ //---> ini function buat ngaapin coba ada ?
    "title": "vaelth notes api backend Golang Web",
    "name":  "valeth",
}

err = tmpl.Execute(w, data)
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
	var message = "Its not you, Its us :("
	w.Write([]byte(message))
}
}

//handler for /hello
func handlerHello(w http.ResponseWriter, r *http.Request) {
    var message = "Hello world!"
    w.Write([]byte(message))
}

//main function
func main() {
    http.HandleFunc("/", handlerwelcome)
    http.HandleFunc("/hello", handlerHello)
		http.Handle("/static/",
        http.StripPrefix("/static/",
            http.FileServer(http.Dir("assets"))))

    var address = "localhost:9000"
    fmt.Printf("server started at %s\n", address)
    err := http.ListenAndServe(address, nil)
    if err != nil {
        fmt.Println(err.Error())
    }
}