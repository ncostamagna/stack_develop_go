package main

import (
	"errors"
	"fmt"

	"gitlab.com/nc_public/sql-client/sqlclient"
)

const (
	queryGetUser = "SELECT id, email FROM users WHERE id=?"
)

var (
	dbClient sqlclient.SqlClient
)

type User struct {
	Id    int64
	Email string
}

func init() {
	sqlclient.StartMockServer()
	var err error
	dbClient, err = sqlclient.Open("maldita", "mierda")

	if err != nil {
		panic(err)
	}
}

func main() {

	user, err := GetUser(123)
	if err == nil {
		panic(err)
	}

	fmt.Println(user.Id)
}

func GetUser(id int64) (*User, error) {
	sqlclient.AddMock(sqlclient.Mock{
		Query: queryGetUser,
		Args:  []interface{}{123},

		Columns: []string{"id", "email"},
		Rows: [][]interface{}{
			{1, "email1"},
			{2, "email2"},
		},
	})
	rows, err := dbClient.Query(queryGetUser, id)

	fmt.Println(rows)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var user User
	for rows.HasNext() {
		if err := rows.Scan(&user.Id, &user.Email); err != nil {
			return nil, err
		}

		fmt.Println(user.Email)
		return &user, nil
	}
	return nil, errors.New("user not found")
}
