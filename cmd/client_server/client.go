package clientserver

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	req_types "sample-choose-ad/cmd/requests_types"
	"sync"
	"time"
)

func makeRequest(url string, body *[]byte, response chan<- []req_types.RespImp, wg *sync.WaitGroup) {
	defer wg.Done()

	var pResp req_types.SuccesResponse

	c := &http.Client{}

	ctx, cls := context.WithTimeout(context.Background(), time.Duration(200*time.Millisecond))
	defer cls()

	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(*body))
	resp, err := c.Do(req)
	// not responding or timeout
	if err != nil {
		log.Println("Error when making request", err)
		return
	}
	defer resp.Body.Close()

	// maybe say smth to client?
	if resp.StatusCode != 200 {
		log.Println("Error: status code", resp.StatusCode)
		return
	}

	b, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(b))
	if err := json.Unmarshal(b, &pResp); err != nil {
		log.Println("Error: response unmarshalling", err)
		return
	}

	// try to convert prices to float

	// for _, imp := range pResp.Imp {
	// 	// log.Printf("%v : %T", imp.PriceStr, imp.PriceStr)
	// 	// imp.Price = imp.PriceStr.(float64)
	// 	// switch imp.PriceStr.(type) {
	// 	// case float64:
	// 	// 	imp.Price = imp.PriceStr.(float64)
	// 	// case string:
	// 	// 	imp.Price, err = strconv.ParseFloat(imp.PriceStr.(string), 64)
	// 	// 	if err != nil {
	// 	// 		log.Println("Pasring price error, ", err)
	// 	// 	}
	// 	// }
	// }

	response <- pResp.Imp
}

// Create requset body based in incoming reqest `ir` and return
// `OutgoingRequest` as []byte from marshaled JSON
func constructPartnersRequestBody(ir *req_types.IncomingRequest) []byte {
	var outReqBody req_types.OutgoingRequest

	var imps []req_types.Imp

	for _, tile := range ir.Tiles {
		minheight := uint(math.Floor(float64(tile.Width) * tile.Ratio))
		imps = append(imps, req_types.Imp{
			Id:        tile.Id,
			Minwidth:  tile.Width,
			Minheight: minheight})
	}

	outReqBody.Id = *ir.Id
	outReqBody.Imp = imps
	outReqBody.Context = ir.Context
	t, _ := json.Marshal(outReqBody)
	return t
}
