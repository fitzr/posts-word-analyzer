package service

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "fmt"
    "io/ioutil"
)

func TestGetStem(t *testing.T) {
    // set up
    var header string
    handler := func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, `{ "text" : "result" }`)

        request, _ := ioutil.ReadAll(r.Body)
        header = string(request)
    }
    server := httptest.NewServer(http.HandlerFunc(handler))
    defer server.Close()
    StemUrl = server.URL
    input := "test"
    expected := "result"
    expectedHeader := "language=english&stemmer=porter&text=test"

    // exercise
    actual := GetStem(input)

    // verify
    if expected != actual {
        t.Errorf("\nexpected: %v\nactual: %v", expected, actual)
    }
    if expectedHeader != header {
        t.Errorf("\nexpected: %v\nactual: %v", expectedHeader, header)
    }
}

/*
func TestGetStemError(t *testing.T) {
    // set up
    handler := func (w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintln(w, "Bad Request")
    }
    server := httptest.NewServer(http.HandlerFunc(handler))
    defer server.Close()
    StemUrl = server.URL
    input := "test"

    // exercise
    GetStem(input)

    // verify log.Fatal()

}
*/
