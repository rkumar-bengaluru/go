package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"regexp"
)

func main() {
	http.HandleFunc("/", handler)
	server := http.Server{}
	server.Addr = ":8080"
	// r := mux.NewRouter()
	// r.HandleFunc("/", handler)
	// server.Handler = r

	// listener, err := net.Listen("tcp", server.Addr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	e := server.ListenAndServe()
	if e != nil {
		log.Fatal(e)
	}
}

func handler(w http.ResponseWriter, req *http.Request) {
	r := regexp.MustCompile("^(.+)@golang.org")
	path := req.URL.Path
	match := r.FindAllStringSubmatch(path, -1)
	//fmt.Println(match, path)
	w.WriteHeader(http.StatusOK)
	if match != nil {
		fmt.Fprintf(w, "Hello from debug %s\n", match[0][0])
		return
	}
	fmt.Fprintln(w, "Hello from debug session")
}
