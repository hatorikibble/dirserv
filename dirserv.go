package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	portPtr := flag.Int("port", 3000, "port number")
	dirPtr := flag.String("dir", ".", "directory")
	flag.Parse()

	listenPar := fmt.Sprintf(":%d", *portPtr)

	fs := http.FileServer(http.Dir(*dirPtr))
	http.Handle("/", fs)

	log.Println(fmt.Sprintf("Serving directory '%s' on port %d", *dirPtr, *portPtr))
	http.ListenAndServe(listenPar, nil)
}
