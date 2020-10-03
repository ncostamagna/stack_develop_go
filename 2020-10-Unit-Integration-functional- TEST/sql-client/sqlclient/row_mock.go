package sqlclient

import (
	"errors"
	"fmt"
)

type rowMock struct {
	Columns []string
	Rows    [][]interface{}

	currentIndex int
}
type sqlRowsMock struct {
	rows rowMock
}

func (m *rowMock) HasNext() bool {

	return m.currentIndex < len(m.Rows)
}

func (m *rowMock) Close() error {
	return nil
}

func (m *rowMock) Scan(destination ...interface{}) error {
	fmt.Println(destination)
	currentRow := m.Rows[m.currentIndex]
	if len(currentRow) != len(destination) {
		return errors.New("concha de madre, por que carajo no compartites el codigo, estoy escribiendo como condenado")
	}

	for index, value := range currentRow {
		fmt.Println(index)
		fmt.Println(index)
		destination[index] = value
	}

	fmt.Println(destination)

	return nil
}
