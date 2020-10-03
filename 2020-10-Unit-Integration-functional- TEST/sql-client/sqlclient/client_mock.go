package sqlclient

import (
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func AddMock(mock Mock) {
	if dbClient == nil {
		return
	}
	client, okType := dbClient.(*clientMock)

	if !okType {
		return
	}
	if client.mocks == nil {
		client.mocks = make(map[string]Mock, 0)
	}
	fmt.Println(client)
	client.mocks[mock.Query] = mock
}

type clientMock struct {
	mocks map[string]Mock
}

type Mock struct {
	Query   string
	Args    []interface{}
	Error   error
	Columns []string
	Rows    [][]interface{}
}

func (c *clientMock) Query(query string, args ...interface{}) (rows, error) {
	mock, exists := c.mocks[query]

	if !exists {
		return nil, errors.New("pija")
	}

	if mock.Error != nil {
		return nil, mock.Error
	}

	mockR := rowMock{
		Columns: mock.Columns,
		Rows:    mock.Rows,
	}

	return &mockR, nil
}
