package task

import (
    "testing"
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
    mock.On("ReadWord").Return("generously").Times(5)
    mock.On("ReadWord").Return("")
    mock.On("WriteStem", "generously", "generous").Times(5)

    // exercise
    MapStem(mock, mock)
}
