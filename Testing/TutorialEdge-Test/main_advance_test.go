package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type AddResult struct {
	x        int
	y        int
	expected int
}

var addResult = []AddResult{
	{1, 2, 3},
}

func TestAdd(t *testing.T) {
	// aca populamos el Add
}

func TestRealFile(t *testing.T) {
	data, err := ioutil.ReadFile("folder/file.data")
	// hacemos las pruebas en base a un archivo
}

func TestHttp(t *testing.T) {
	//
	handler := func(w http.ResponseWriter, r *http.Request) {
		// here we write our expected response, in this case, we return a
		// JSON string which is typical when dealing with REST APIs
		io.WriteString(w, "{ \"status\": \"expected service response\"}")
	}

	req := httptest.NewRequest("GET", "https://tutorialedge.net", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}
