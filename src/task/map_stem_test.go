package task

import (
    "../service"
    "testing"
    "fmt"
    "net/http"
    "net/http/httptest"
    "github.com/stretchr/testify/mock"
)

type stemMock struct {
    mock.Mock
}

func (m *stemMock) ReadWord() string {
    arg := m.Called()
    return arg.String(0)
}

func (m *stemMock) WriteStem(word, stem string) {
    m.Called(word, stem)
}

func TestMapStem(t *testing.T) {

    // reader and writer mock
    mock := new(stemMock)
    mock.On("ReadWord").Return("word").Times(5)
    mock.On("ReadWord").Return("")
    mock.On("WriteStem", "word", "stem").Times(5)

    // http mock
    handler := func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, `{ "text" : "stem" }`)
    }
    server := httptest.NewServer(http.HandlerFunc(handler))
    defer server.Close()
    service.StemUrl = server.URL

    // exercise
    MapStem(mock, mock)
}
