package go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func HelloHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(rw, "Hello World")
}

func TestHttp(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
	rec := httptest.NewRecorder()

	HelloHandler(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func SayHello(rw http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(rw, "Hewwo")
	} else {
		fmt.Fprintf(rw, "Hewwo %s", name)
	}
}

func TestQueryParams(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Dimas", nil)
	rec := httptest.NewRecorder()

	SayHello(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func MultipleParams(rw http.ResponseWriter, r *http.Request) {
	firstname := r.URL.Query().Get("firstname")
	lastname := r.URL.Query().Get("lastname")
	fmt.Fprintf(rw, "Hewwo %s %s", firstname, lastname)
}

func TestMultipleParams(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?firstname=Dimas&lastname=Eka", nil)
	rec := httptest.NewRecorder()

	MultipleParams(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func MultipleValue(rw http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprintln(rw, strings.Join(names, " "))
}

func TestMultipleValue(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Dimas&name=Eka&name=Adinandra", nil)
	rec := httptest.NewRecorder()

	MultipleValue(rec, req)

	response := rec.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}
