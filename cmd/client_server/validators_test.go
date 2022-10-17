package clientserver

import (
	"testing"
)

// Wants: 10.10.10.10:5050
// Gets : localhost:5050
func TestIPAddresFormat_DomainName(t *testing.T) {
	addres := "localhost:5050"
	_, _, e := ParsePartnersAddress(addres)
	if e == nil {
		t.Error("Must be an error, when parsing ", addres)
	}
	t.Log(e)
}

// Wants: 10.10.10.10:5050
// Gets : 10.10.10.10:
func TestIPAddresFormat_OnlyIpAndColon(t *testing.T) {
	addres := "10.10.10.10:"
	_, _, e := ParsePartnersAddress(addres)
	if e == nil {
		t.Error("Must be an error, when parsing ", addres)
	}
	t.Log(e)
}

// Wants: 10.10.10.10:5050
// Gets : 10.10.10.10
func TestIPAddresFormat_OnlyIp(t *testing.T) {
	addres := "10.10.10.10"
	_, _, e := ParsePartnersAddress(addres)
	if e == nil {
		t.Error("Must be an error, when parsing ", addres)
	}
	t.Log(e)
}

// Wants: 10.10.10.10:5050
// Gets : 10.10.10.10:65537
func TestIPAddresFormat_IncorrectPortValue_TooBig(t *testing.T) {
	addres := "10.10.10.10:65537"
	_, _, e := ParsePartnersAddress(addres)
	if e == nil {
		t.Error("Must be an error, when parsing ", addres)
	}
	t.Log(e)
}

// Wants: 10.10.10.10:5050
// Gets : 10.10.10.10:1000
func TestIPAddresFormat_IncorrectPortValue_TooSmall(t *testing.T) {
	addres := "10.10.10.10:1000"
	_, _, e := ParsePartnersAddress(addres)
	if e == nil {
		t.Error("Must be an error, when parsing ", addres)
	}
	t.Log(e)
}

// Wants: 10.10.10.10:5050
// Gets : 10.10.10.10:as
func TestIPAddresFormat_IncorrectPortValue_NotANumber(t *testing.T) {
	addres := "10.10.10.10:as"
	_, _, e := ParsePartnersAddress(addres)
	if e == nil {
		t.Error("Must be an error, when parsing ", addres)
	}
	t.Log(e)
}

// Wants: 10.10.10.10:5050
// Gets : 10.10.10.10:5050/bid_request
func TestIPAddresFormat_AddressWithEndpoint(t *testing.T) {
	addres := "10.10.10.10:5050/bid_request"
	_, _, e := ParsePartnersAddress(addres)
	if e == nil {
		t.Error("Must be an error, when parsing ", addres)
	}
	t.Log(e)
}
