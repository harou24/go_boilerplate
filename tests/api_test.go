package tests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/harou24/go_boilerplate/api"
)

func TestServe(t *testing.T) {
	// Create a new Manager instance
	manager := api.NewManager()
	// Start the server on a new goroutine
	go manager.Serve("5000")
	// Wait for the server to start
	time.Sleep(1 * time.Second)

	// Create a new http client
	client := &http.Client{}
	// Make a GET request to the / endpoint
	resp, err := client.Get("http://localhost:5000/")
	if err != nil {
		t.Errorf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check that the response status is 200 OK
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK but got %v", resp.Status)
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body: %v", err)
	}
	fmt.Println(string(body))

	// Check that the response body contains the expected message
	if !strings.Contains(string(body), "Hello World 🖖") {
		t.Errorf("Expected response body to contain 'Hello World 🖖' but got %s", string(body))
	}
}
