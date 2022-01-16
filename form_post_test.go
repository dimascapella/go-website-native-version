package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(rw http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	firstname := r.PostForm.Get("firstname")
	lastname := r.PostForm.Get("lastname")
	fmt.Fprintf(rw, "Hewoo %s %s", firstname, lastname)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstname=Dimas&lastname=EkaAdinandra")
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requestBody)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	FormPost(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
