package task

import (
    "testing"
    "github.com/stretchr/testify/mock"
)

type stemMock struct {
    mock.Mock
}

func (m *stemMock) Read() string {
    arg := m.Called()
    return arg.String(0)
}

func (m *stemMock) Write(args ...interface{}) {
    m.Called(args...)
}

func TestMapStem(t *testing.T) {

    // reader and writer mock
    mock := new(stemMock)
    mock.On("Read").Return("generously").Times(5)
    mock.On("Read").Return("")
    mock.On("Write", "generously", "generous").Times(5)

    // exercise
    MapStem(mock, mock)
}
