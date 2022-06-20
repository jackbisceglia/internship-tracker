package crud

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	Id string
	CreatedAt string
	EmailAddress string
	PreferenceList string
}

func UserCrud(db *sql.DB) (func(string) []User, func(string, []string) bool) {
	GetUsersByList := func(preferenceList string) []User {
		var user User
		db_values := []interface{}{}
		SQL_STATEMENT := `
			SELECT * FROM users
		`

		if preferenceList != "BOTH" {
			SQL_STATEMENT = fmt.Sprintf("%s WHERE preferenceList = $1 or preferenceList = $2", SQL_STATEMENT)
			db_values = append(db_values, "BOTH", preferenceList)
		}


		rows, err := db.Query(SQL_STATEMENT, db_values...)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
	
		var users []User
		for rows.Next() {
			err := rows.Scan(&user.Id, &user.CreatedAt, &user.EmailAddress, &user.PreferenceList)
			if err != nil {
				log.Fatal(err)
			}
	
			users = append(users, user)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		return users
	}

	InsertUser := func(emailAddress string, listPreferences []string) bool {
		var preferenceString string

		if len(listPreferences) > 1 {
			preferenceString = "BOTH"
		} else {
			preferenceString = listPreferences[0]
		}

		SQL_STATEMENT := `
			INSERT INTO users (emailaddress, preferencelist)
			VALUES ($1, $2)
		`

		_, err := db.Exec(SQL_STATEMENT, emailAddress, preferenceString)
		if err != nil {
			fmt.Printf("%v", err)
			return false
		}

		return true
	}

	return GetUsersByList, InsertUser
}
