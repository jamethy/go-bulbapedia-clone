package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

var (
	listen = flag.String("listen", ":8080", "listen address")
	dir    = flag.String("dir", ".", "directory to serve")
)

type MyHandler struct {
	handler http.Handler
}

// hacky method to route to / if page not found - which is mostly always true for a SPA
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	log.Println("received ", upath)
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	fileName := path.Clean(upath)
	if _, err := os.Open("./" + fileName); err != nil {
		r.URL.Path = "/"
		log.Println("updated to ", r.URL.Path)
	}
	h.handler.ServeHTTP(w, r)
}

func main() {
	flag.Parse()
	log.Printf("listening on %q...", *listen)
	log.Fatal(http.ListenAndServe(*listen, &MyHandler{handler: http.FileServer(http.Dir(*dir))}))
}
