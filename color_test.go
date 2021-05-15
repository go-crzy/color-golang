package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_webServer(t *testing.T) {
	handler := &webServer{}
	server := httptest.NewServer(handler)
	client := server.Client()

	request, _ := http.NewRequest("Get", server.URL, nil)
	response, err := client.Do(request)
	if err != nil {
		t.Error("Should succeed", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf(
			"Status Code should be 200, current: %d",
			response.StatusCode,
		)
	}
}

