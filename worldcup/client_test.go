package worldcup

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// setup sets up a test HTTP server along with a 'worldcup.Client' that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, server *httptest.Server, teardown func()) {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient()
	url, _ := url.Parse(server.URL + "/")
	client.baseURL = url

	return client, mux, server, server.Close
}

// fixture reads a 'testdata' file and returns it as a string
func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func TestNewClient(t *testing.T) {
	client := NewClient()
	baseURL, _ := url.Parse(apiURL)

	assert.IsType(t, new(Client), client)
	assert.Equal(t, client.baseURL, baseURL)
	assert.Equal(t, client.httpClient, http.DefaultClient)
}
