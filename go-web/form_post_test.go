package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// representasi request form user

// handler
func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

// use case
func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Noby&last_name=Nobygon")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder() // empty http response

	FormPost(recorder, request) // manipulate http reponse content

	response := recorder.Result() // convert the result into *http.Reponse type
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
