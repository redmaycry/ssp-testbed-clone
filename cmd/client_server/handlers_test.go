package clientserver

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	customtypes "sample-choose-ad/cmd/custom_types"
	req_types "sample-choose-ad/cmd/requests_types"
	"testing"
)

func TestGetRequestWithEmptyBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/placements/request", nil)
	w := httptest.NewRecorder()
	_, _ = req, w
	a := handleRequest([]customtypes.PartnersAddress{{Ip: "127.0.0.1", Port: 5050}})
	a(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expects code 400, got %v", res.StatusCode)
	}

}

func TestPostRequestWithEmptyBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/placements/request", nil)
	w := httptest.NewRecorder()
	_, _ = req, w
	a := handleRequest([]customtypes.PartnersAddress{{Ip: "127.0.0.1", Port: 5050}})
	a(w, req)
	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expects code 400, got %v", res.StatusCode)
	}
}

func TestPostRequestWithRightBody(t *testing.T) {
	body_json := `{
  "id": "123",
  "tiles": [
    {
      "id": 123,
      "width": 122,
      "ratio": 1.5
    }
  ],
  "context": {
    "ip": "192.168.1.1",
    "user_agent": "curl"
  }
}`

	req := httptest.NewRequest(http.MethodPost, "/placements/request", bytes.NewBuffer([]byte(body_json)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	_, _ = req, w
	a := handleRequest([]customtypes.PartnersAddress{{Ip: "127.0.0.1", Port: 5059}})
	a(w, req)
	res := w.Result()
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expects code 200, got %v", res.StatusCode)
	}

	var d req_types.SuccesResponse
	if json.Unmarshal(data, &d) != nil {
		t.Log("Error parsing json response")
	}

	if d.Imp[0].Title != "bestoption" {
		t.Errorf("Wants title `bestoption`, got %v", d.Imp[0].Title)
	}

}
