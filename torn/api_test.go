package torn

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This round trip approach was taken from http://hassansin.github.io/Unit-Testing-http-client-in-Go
// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

//NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
func TestGetCooldown(t *testing.T) {
	body := `{"chain":{"current":0,"max":10,"timeout":0,"modifier":1,"cooldown":150,"start":0}}`
	r := ioutil.NopCloser(strings.NewReader(body))
	fakeClient := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       r,
		}
	})

	cd, err := GetChainCooldown(fakeClient, "")

	assert.NoError(t, err)
	assert.Equal(t, 150, cd)
}