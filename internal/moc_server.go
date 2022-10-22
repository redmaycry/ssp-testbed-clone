package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.ReadFile("internal/json/valid_response.json")
	if err != nil {
		log.Fatalln(err)
	}

	addr := flag.String("l", "", "-l 127.0.0.1:5059")
	flag.Parse()
	if *addr == "" {
		log.Fatalln("Error: listening address is required!")
	}

	http.HandleFunc("/bid_request", func(w http.ResponseWriter, r *http.Request) {
		inc, _ := ioutil.ReadAll(r.Body)
		log.Println(string(inc))
		w.Header().Add("Content-Type", "application/json")
		w.Write(file)
	})

	// endpoint: /exit
	// Terminate server with code 0.
	http.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		os.Exit(0)
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}
