package main

import (
	"github.com/gobestsdk/gobase/log"
	"github.com/timeloveboy/apiman/htmlpart"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	root     = ""
	port     = "8009"
	notexist = "not exist"
)

func staticDirHandler(mux *http.ServeMux, prefix string) {
	mux.HandleFunc(prefix,
		func(w http.ResponseWriter, r *http.Request) {
			log.Info(log.Fields{"path": r.URL.Path})
			root, _ = os.Getwd()
			f := root + r.URL.Path
			fi, err := os.Lstat(f)
			if err != nil {
				log.Fatal(log.Fields{"err": err})
				w.Write([]byte(notexist))
				return
			}

			switch mode := fi.Mode(); {
			case mode.IsRegular():
				bs, _ := ioutil.ReadFile(f)
				result := htmlpart.Render(r.URL.Path, string(bs))
				w.Write([]byte(result))
			default:
				log.Info(log.Fields{"err": err})
				w.Write([]byte("notexist"))
			}
		})
}

func main() {
	log.Setlogfile("opengw.log")

	var mux = http.NewServeMux()
	staticDirHandler(mux, "/")
	log.Info(log.Fields{"http.ListenAndServe": port})

	err := http.ListenAndServe(":"+port, mux)

	if err != nil {
		log.Fatal(log.Fields{"http.ListenAndServe:": err.Error()})
	}
}
