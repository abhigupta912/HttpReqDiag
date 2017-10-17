package main

import (
	"HttpReqDiag/server"
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var port = flag.Int("port", 9000, "Port Number")
	flag.Parse()

	log.Printf("Starting server on port: %d\n", *port)
	http.Handle("/", server.NewReqParser(*port))
	http.Handle("/dump", server.NewReqDumper(*port))
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*port), nil))
}
