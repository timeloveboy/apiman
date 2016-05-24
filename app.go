package main

import (
	"log"
	"net/http"
	"os"
)

var (
	root = ""
	data = "/"
	port = "8001"
)

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix,
		func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.URL.Path)
			file := root + r.URL.Path
			log.Println(file)
			http.ServeFile(w, r, file)
		})
}
func main() {
	root, _ = os.Getwd()
	var mux = http.NewServeMux()
	staticDirHandler(mux, "/", root+data, 0)

	log.Println("http.ListenAndServe(:" + port + ")")
	err := http.ListenAndServe(":"+port, mux)

	if err != nil {
		log.Fatal("http.ListenAndServe:", err.Error())
	}
}
