package clientserver

import (
	"fmt"
	"log"
	"net/http"
	customtypes "sample-choose-ad/cmd/custom_types"
)

func StartServer(port string, partners []customtypes.PartnersAddress) {

	http.HandleFunc("/placements/request", handleRequest(partners))
	// http.HandleFunc("/placements/request", decorate(test2))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}
