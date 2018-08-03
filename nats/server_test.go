package nats

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/horaddrim/natsctl/utils"
)

// TestServerMarshalling is just a illusion
func TestServerMarshalling(t *testing.T) {
	client := utils.NewHTTPClient()
	server := new(Server)

	resp, err := client.Get("http://40.121.32.164:8222/connz")

	if err != nil {
		t.Fail()
	}
	buff, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(buff, server)

	if server.ID != "vZLdOqiDgE6M0aZGLKbxpz" {
		t.Fatalf("Parsing failed.")
	}
}
