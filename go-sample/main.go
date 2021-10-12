package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const INDEX = `<!DOCTYPE html>
<html>
  <head>
    <title>Go Sample</title>
    <link rel='stylesheet' href='/stylesheets/style.css'/>
  </head>
  <body>
    <h1>Go on AWS App Runner with Cloud Native Buildpacks</h1>
    <p>This application is running as a container on AWS App Runner and was built using Cloud Native Buildpacks!</p>
  </body>
</html>
`

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, INDEX)
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
