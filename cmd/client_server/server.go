package clientserver

import (
	"fmt"
	"log"
	"net/http"
	customtypes "sample-choose-ad/cmd/custom_types"
	"time"
)

const MAX_TIME_PER_REQUEST = time.Duration(250 * time.Millisecond)

type customHandler struct {
	Parners []customtypes.PartnersAddress
	// context?
}

func (c *customHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	base := r.URL.Path

	switch base {
	case "/placements/request":

	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func newCustomHandler() *customHandler {
	return &customHandler{}
}

func StartServer(port string, partners []customtypes.PartnersAddress) {
	// mux := http.NewServeMux()
	// mux.HandleFunc("/placements/request", handleRequest(partners))

	// s := &http.Server{
	// 	ReadTimeout:  time.Duration(time.Millisecond * 20),
	// 	WriteTimeout: time.Duration(time.Millisecond * 20),
	// 	Handler:      newCustomHandler(),
	// }
	// s.ListenAndServe()
	// h := http.HandleFunc("/placements/request", handleRequest(partners))
	h := http.TimeoutHandler(handleRequest(partners), MAX_TIME_PER_REQUEST, "")
	http.Handle("/placements/request", h)
	// http.Handle("/placements/request", handleRequest(partners))
	// http.HandleFunc("/placements/request", decorate(test2))

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}
