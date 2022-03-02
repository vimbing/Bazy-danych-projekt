package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetData() []User {
	db, err := sql.Open("mysql", "program:program@tcp(127.0.0.1:3306)/uzytkownicy")

	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT * FROM `uzytkownicy`")

	if err != nil {
		panic(err.Error())
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var allUsers []User

	for rows.Next() {

		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		allUsers = append(allUsers, User{Login: string(values[0]), Password: string(values[1])})
	}

	if err = rows.Err(); err != nil {
		panic(err.Error())
	}

	return allUsers
}
