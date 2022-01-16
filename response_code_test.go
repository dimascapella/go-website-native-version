package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(rw, "Hewwo")
	} else {
		rw.WriteHeader(200)
		fmt.Fprintf(rw, "Hewwo %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	rec := httptest.NewRecorder()

	ResponseCode(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
	fmt.Println(response.Status)
}
