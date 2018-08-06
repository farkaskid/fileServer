package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type fileHandler struct {
	fsHandler http.Handler
	cwd       string
}

func (handler fileHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	log.Println(req.RemoteAddr, "accessed", handler.cwd+req.URL.String(), "using", req.Method)
	handler.fsHandler.ServeHTTP(writer, req)
}

func main() {
	port := flag.String("port", "4000", "port number")
	path := flag.String("path", "", "directory to serve")

	flag.Parse()

	if *path == "" {
		var err error
		*path, err = os.Getwd()

		if err != nil {
			log.Fatalln("Failed to get the current working directory")
		}
	}

	fileServer := http.FileServer(http.Dir(*path))

	http.Handle("/", fileHandler{fileServer, *path})

	err := http.ListenAndServe(":"+*port, nil)

	if err != nil {
		log.Fatalln("Failed to start server at port", *port)
	}

	log.Println("FileServer started at port", *port)
}
