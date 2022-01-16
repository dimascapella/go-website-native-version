package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(rw http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(rw, contentType)
}

func TestRequestHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	req.Header.Add("Content-Type", "application/json")

	rec := httptest.NewRecorder()

	RequestHeader(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func ResponseHeader(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("X-Powered-By", "Dimas Capella")
	fmt.Fprint(rw, "OK")
}

func TestResponseHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	req.Header.Add("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	ResponseHeader(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
	fmt.Println(response.Header.Get("x-powered-by"))
}
