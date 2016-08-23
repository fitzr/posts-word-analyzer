package task

import (
    "github.com/stretchr/testify/mock"
    "testing"
    "net/http"
    "fmt"
    "net/http/httptest"
    "../service"
)

type attrMock struct {
    mock.Mock
}

func (m *attrMock) Read() string {
    arg := m.Called()
    return arg.String(0)
}

func (m *attrMock) Write(args ...interface{}) {
    m.Called(args...)
}

func TestMapAttr(t *testing.T) {
    // reader and writer mock
    mock := new(attrMock)
    mock.On("Read").Return("word").Times(32)
    mock.On("Read").Return("")
    mock.On("Write", "word", 5.11, "noun").Times(32)

    // http mock
    handler := func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, `{"frequency":5.11,"results":[{"partOfSpeech":"noun"}]}`)
    }
    server := httptest.NewServer(http.HandlerFunc(handler))
    defer server.Close()
    fmt.Println(server.URL)
    service.StemUrl = server.URL

    // exercise
    MapAttr(mock, mock)
}
