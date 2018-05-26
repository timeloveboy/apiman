package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gobestsdk/gobase/log"
)

var (
	root = ""
	data = "/"
	port = "8001"
	notexist="not exist"
)

func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix,
		func(w http.ResponseWriter, r *http.Request) {
			log.Info(log.Fields{"path":r.URL.Path})
			root, _ = os.Getwd()
			f := root + r.URL.Path
			fi, err := os.Lstat(f)
			if err != nil {
				log.Fatal(log.Fields{"err":err,})
				w.Write([]byte(notexist))
			}

			switch mode := fi.Mode(); {
				case mode.IsRegular():
					htmlpart(w, r)
				default  :
					log.Info(log.Fields{"err":err})
					w.Write([]byte("notexist"))
			}
		})
}

func htmlpart(w http.ResponseWriter, r *http.Request){
	

}
func main() {
	log.Setlogfile("opengw.log")
	
	var mux = http.NewServeMux()
	staticDirHandler(mux, "/", root+data, 0)
	log.Info(log.Fields{"http.ListenAndServe":port})
 
	err := http.ListenAndServe(":"+port, mux)

	if err != nil {
		log.Fatal(log.Fields{"http.ListenAndServe:": err.Error()})
	}
}
