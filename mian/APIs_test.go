package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetAllTrans(t *testing.T) {
	req, err := http.NewRequest("GET", "/transactions/", nil)
	if err != nil {
		t.Fatal(err)
	}
	reqTest := httptest.NewRecorder()
	handler := http.HandlerFunc(getAll)
	handler.ServeHTTP(reqTest, req)
	if status := reqTest.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `[{"Id":"6c59fbd9f3424059a70430b8e844b078","Amount":50000000,"Currency":"CLP","CreatedAt":"2022-08-02T11:11:55.5465019Z"},{"Id":"2b7133bfcfb844179f7e19f18c77725a","Amount":36457,"Currency":"MXN","CreatedAt":"2022-08-02T11:11:55.5465019Z"},{"Id":"5135421cf0ff47959fecd8670989807b","Amount":1151.12,"Currency":"USD","CreatedAt":"2022-08-02T11:11:55.5465019Z"},{"Id":"c6854925270841559069812a64e288e7","Amount":9965000.4,"Currency":"COP","CreatedAt":"2022-08-02T11:11:55.5465019Z"},{"Id":"27313a7f2b5a4281aec03b4b68090118","Amount":12500,"Currency":"BRA","CreatedAt":"2022-08-02T11:11:55.5465019Z"}]`
	if strings.Compare(reqTest.Body.String(), expected) == 0 {
		t.Errorf("handler returned unexpected body: got %v want %v",
			reqTest.Body.String(), expected)
	}
}

func TestAddTrans(t *testing.T) {

	var jsonStr = []byte(`{"ID":4,"Amount":365.78,"Currency":"CLP"`)

	req, err := http.NewRequest("POST", "/transactions", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	reqTest := httptest.NewRecorder()
	handler := http.HandlerFunc(addTrans)
	handler.ServeHTTP(reqTest, req)
	if status := reqTest.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `Transaction was added successfully`
	if reqTest.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			reqTest.Body.String(), expected)
	}
}

func TestGetTransNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/transactions/", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("Id", "49b6e6d863f547db805d1eb208e3a44a")
	req.URL.RawQuery = q.Encode()
	reqTest := httptest.NewRecorder()
	handler := http.HandlerFunc(get)
	handler.ServeHTTP(reqTest, req)
	if status := reqTest.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}
