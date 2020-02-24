package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"skoltech/pkg/app"
	"skoltech/pkg/devices"
	"testing"
)

type serviceMock struct {
}

func (s *serviceMock) CreateDevice() *devices.Device {
	return devices.NewDevice()
}

func (s *serviceMock) Save(d *devices.Device) error {
	return nil
}

func TestNotPostRequest(t *testing.T) {

	sMock := &serviceMock{}
	server := app.NewServer(sMock)

	srv := httptest.NewServer(server.Handlers())
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/", srv.URL))

	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status has to be %v", http.StatusMethodNotAllowed)
	}
	defer res.Body.Close()
}

func TestNotJSONRequest(t *testing.T) {
	sMock := &serviceMock{}
	server := app.NewServer(sMock)

	srv := httptest.NewServer(server.Handlers())
	defer srv.Close()

	res, err := http.Post(fmt.Sprintf("%s/", srv.URL), "html/text", nil)
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != http.StatusUnsupportedMediaType {
		t.Errorf("Status has to be %v", http.StatusUnsupportedMediaType)
	}
	defer res.Body.Close()
}

func TestDeviceRequest(t *testing.T) {
	sMock := &serviceMock{}
	server := app.NewServer(sMock)

	srv := httptest.NewServer(server.Handlers())
	defer srv.Close()
	res, err := http.Post(
		fmt.Sprintf("%s/", srv.URL),
		"application/json",
		bytes.NewBuffer(
			[]byte(`{"ap_id" : "A8-F9-4B-B6-87-FF", "version" : "1.0", "probe_requests" : [{"mac" : "88-1D-FC-DF-6F-C1","timestamp" : "1579782767"}]}`),
		),
	)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status has to be %v", http.StatusOK)
	}
	defer res.Body.Close()
}

func TestDeviceBadRequest(t *testing.T) {
	sMock := &serviceMock{}
	server := app.NewServer(sMock)

	srv := httptest.NewServer(server.Handlers())
	defer srv.Close()

	cases := []struct {
		str string
	}{
		{"!"},
		{"{,}"},
	}

	for _, tc := range cases {
		res, err := http.Post(fmt.Sprintf("%s/", srv.URL), "application/json", bytes.NewBuffer([]byte(tc.str)))
		if err != nil {
			t.Fatal(err)
		}
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("Status has to be %v", http.StatusBadRequest)
		}
		defer res.Body.Close()
	}
}
