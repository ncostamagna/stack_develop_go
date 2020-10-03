package sqlclient

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	goEnv      = "GO_ENVIRONMENT"
	production = "production"
)

var (
	isMocked bool
	dbClient SqlClient
)

type client struct {
	db *sql.DB
}

type SqlClient interface {
	Query(query string, args ...interface{}) (rows, error)
}

func StartMockServer() {
	isMocked = true
}

func StopMockServer() {
	isMocked = false
}

func isProducction() bool {
	return os.Getenv(goEnv) == production
}

func Open(driveName, dataSourceName string) (SqlClient, error) {

	if isMocked && !isProducction() {
		fmt.Println("Entra")
		dbClient = &clientMock{}
		return dbClient, nil
	}
	if driveName == "" {
		return nil, errors.New("Invalid client drive name")
	}

	db, err := sql.Open(driveName, dataSourceName)

	if err != nil {
		return nil, err
	}

	dbClient = &client{
		db: db,
	}

	return dbClient, nil
}

func (c *client) Query(query string, args ...interface{}) (rows, error) {

	returnedRows, err := c.db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	result := sqlRows{
		rows: returnedRows,
	}

	return &result, nil
}
