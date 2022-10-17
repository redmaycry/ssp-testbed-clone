package clientserver

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	customtypes "sample-choose-ad/cmd/custom_types"
	req_types "sample-choose-ad/cmd/requests_types"
	"sort"
	"time"
)

func sendRequest(url string, body *[]byte) (req_types.SuccesResponse, error) {
	var pResp req_types.SuccesResponse

	c := &http.Client{
		Timeout: 200 * time.Millisecond,
	}

	resp, err := c.Post(url, "application/json", bytes.NewReader(*body))

	if err != nil {
		eText := fmt.Sprintf("Error: partner %v not responding", url)
		return pResp, errors.New(eText)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 204 {
		return pResp, errors.New("No content")
	}

	b, _ := ioutil.ReadAll(resp.Body)

	if json.Unmarshal(b, &pResp) != nil {
		log.Println(err)
	}

	return pResp, nil
}

// Create requset body based in incoming reqest `ir` and return
// `OutgoingRequest` as []byte from marshaled JSON
func constructPartnersRequestBody(ir *req_types.IncomingRequest) []byte {
	var outReqBody req_types.OutgoingRequest

	var imps []req_types.Imp

	// WARN: uint and float multiplication may cause problems
	for _, tile := range ir.Tiles {
		imps = append(imps, req_types.Imp{
			Id:        tile.Id,
			Minwidth:  tile.Width,
			Minheight: uint(math.Floor(float64(tile.Width * uint(tile.Ratio))))})
	}

	outReqBody.Id = *ir.Id
	outReqBody.Imp = imps
	outReqBody.Context = ir.Context
	t, _ := json.Marshal(outReqBody)
	return t
}

// map[imp.id]map[imp.id.price]
type PartnersResponses map[uint]map[float64]req_types.RespImp

// Make request for each partner and returns
func makePartnersRequests(partners []customtypes.PartnersAddress, ir *req_types.IncomingRequest) {
	p_body := constructPartnersRequestBody(ir)

	// Two data structures:
	// partnersRespones for getting price with O(1) complexity
	// []prices as slice of actual prices
	partnersRespones := make(map[uint]map[float64]req_types.RespImp)
	prices := make(map[uint][]float64)

	for _, p := range partners {
		url := fmt.Sprintf("http://%v:%v/%v", p.Ip, p.Port, PARTNER_ENDPOINT)

		re, err := sendRequest(url, &p_body)

		if err != nil {
			log.Println(err)
			continue
		}
		// append only successful responses
		for _, r := range re.Imp {
			if partnersRespones[r.Id] == nil {
				partnersRespones[r.Id] = make(map[float64]req_types.RespImp)
			}
			partnersRespones[r.Id][r.Price] = r
			prices[r.Id] = append(prices[r.Id], r.Price)
		}

	}

	if len(partnersRespones) == 0 {
		log.Println("Error: no responses from partners.")
		return
	}

	// Sorting prices, now biggest price at index len-1
	for _, p := range prices {
		sort.Float64s(p)
	}
}
