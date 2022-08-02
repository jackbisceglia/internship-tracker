package crud

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
)

type PostingData struct {
	Company string `json:"company"`
	Location string `json:"location"`
	Notes string `json:"notes"`
	IsIntern bool `json:"isIntern"`
	Url string `json:"url"`
	Id string
	CreatedAt string
}

func PostingsCrud(db *sql.DB) (func() []PostingData, func([]PostingData) bool, func(postingUrl string) error ){
	GetPostings := func() []PostingData {
		var posting PostingData
		SQL_STATEMENT := `
			SELECT * FROM postings
		`

		rows, err := db.Query(SQL_STATEMENT)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
	
		var postings []PostingData
		for rows.Next() {
			err := rows.Scan(&posting.Id, &posting.Company, &posting.Location, &posting.Notes, &posting.CreatedAt, &posting.IsIntern, &posting.Url)
			if err != nil {
				log.Fatal(err)
			}
	
			postings = append(postings, posting)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		return postings
	}

	InsertPosting := func(posts []PostingData) bool {
		SQL_STATEMENT := `
			INSERT INTO postings (company, location, notes, isintern, url)
			VALUES`

		numFields := 5
		values := []interface{}{}
		for i, post := range posts {
			values = append(values, post.Company, post.Location, post.Notes, post.IsIntern, post.Url)

			n := numFields * i

			queryFields := "("
			for j := 0; j < numFields; j++ {
				queryFields = fmt.Sprintf("%s$%s,", queryFields, strconv.Itoa(n + j + 1))
			}
			queryFields = queryFields[:len(queryFields)-1] + `),`
			SQL_STATEMENT += queryFields
		}

		SQL_STATEMENT = SQL_STATEMENT[:len(SQL_STATEMENT)-1]

		_, err := db.Exec(SQL_STATEMENT, values...)
		if err != nil {
			fmt.Printf("%v", err)
			return false
		}

		return true
	}
	
	DeletePosting := func(postingUrl string) error {

		SQL_STATEMENT := `
			DELETE FROM postings WHERE url = $1;
		`

		result, err := db.Exec(SQL_STATEMENT, postingUrl)
		if err != nil {
			fmt.Printf("%v", err)
			return errors.New("There was an error deleting")
		}

		numDeleted, err := result.RowsAffected();

		if err != nil {
			fmt.Printf("%v", err)
			return errors.New("There was an error deleting")
		}

		if numDeleted == 0 {
			return errors.New("No posting found with that url")
		}

		return nil
	}

	return GetPostings, InsertPosting, DeletePosting
}