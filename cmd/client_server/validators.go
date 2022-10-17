package clientserver

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	MAX_PORT_NUM = 65535
	MIN_PORT_NUM = 1024
)

// Returns false if ipv4 `correct`.
func wrongIPAddresFormat(ipv4 string) bool {
	re, err := regexp.Compile(`^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`)
	if err != nil {
		log.Println(err)
	}
	return !re.Match([]byte(ipv4))
}

func throwHTTPError(err_text string, code int, w *http.ResponseWriter) error {
	http.Error(*w, err_text, code)
	eText := fmt.Sprintf("Error: %d %vr", code, err_text)
	return errors.New(eText)
}

// Wait string in format "10.10.10.10:8080", where `10.10.10.10` IPv4,
// and `8080` port. If ip or port has wrong format, returns error.
func ParsePartnersAddress(ipAndPort string) (string, int64, error) {
	var err error
	var ip string
	var port int64

	iap := strings.Split(ipAndPort, ":")

	if len(iap) != 2 {
		err = errors.New(fmt.Sprintf("Wrong partners 'ip:port' format: %v", ipAndPort))
		return ip, port, err
	}

	ip = iap[0]
	if wrongIPAddresFormat(ip) {
		err = errors.New(fmt.Sprintf("Wrong ip address format in partner ip: %v", ip))
	}

	port, e := strconv.ParseInt(iap[1], 10, 64)
	if e != nil {
		err = errors.New(fmt.Sprintf("Wrong port format in partner ip: %v", e))
		return ip, port, err
	}

	if port > MAX_PORT_NUM {
		err = errors.New(fmt.Sprintf("Wrong port in partner ip: grater than %v", MAX_PORT_NUM))
	}

	if port < MIN_PORT_NUM {
		err = errors.New(fmt.Sprintf("Wrong port in partner ip: %v lower than %v", port, MIN_PORT_NUM))
	}

	return ip, port, err
}
