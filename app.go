package main

import (
	"flag"
	"github.com/gobestsdk/gobase/httpserver"
	"github.com/gobestsdk/gobase/log"
	"github.com/timeloveboy/apiman/htmlpart"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	server   = httpserver.New()
	root     = ""
	port     = 80
	notexist = "not exist"
)

func init() {

	flag.StringVar(&root, "root", "", "--root=/myweb")
	flag.IntVar(&port, "port", 80, "--port=8080")
}

func main() {
	log.Setlogfile("apiman.log")

	flag.Parse()
	defer func() {
		if error := recover(); error != nil {
			log.Fatal(log.Fields{"panic": error})
			exit(-1)
		}
	}()

	go func() {
		singals := make(chan os.Signal)
		signal.Notify(singals,
			syscall.SIGTERM,
			syscall.SIGINT,
			syscall.SIGKILL,
			syscall.SIGHUP,
			syscall.SIGQUIT,
		)
		<-singals
		exit(0)
	}()

	if root == "" {
		root, _ = os.Getwd()
	}

	server.Mux.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			log.Info(log.Fields{"path": r.URL.Path})

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
				result := htmlpart.Render(root, r.URL.Path, string(bs))
				w.Write([]byte(result))
			default:
				log.Info(log.Fields{"err": err})
				w.Write([]byte("notexist"))
			}
		})
	server.SetPort(port)
	server.Run()

}
func exit(status int) {
	log.Info(log.Fields{"app": status})
	server.Stop()
	os.Exit(status)
}
