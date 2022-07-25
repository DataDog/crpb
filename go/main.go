package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start(tracer.WithDebugMode(true))
	defer tracer.Stop()
	mux := httptrace.NewServeMux()
	mux.HandleFunc("/", handle)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}
