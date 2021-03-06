package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackbisceglia/internship-tracker/crud"
	"github.com/jackbisceglia/internship-tracker/util"
)

type PostResponse struct {
	InternPosts []crud.PostingData
	NewGradPosts []crud.PostingData
}

func PostingRoutes(router *mux.Router, db *sql.DB) {
	HandleMultiplePostingRoutes := util.RouterUtils(router)
	GetPostings, InsertPosting := crud.PostingsCrud(db)

	getPostingsHandler := func(w http.ResponseWriter, r *http.Request) {
		postings := GetPostings()
		internList := make([]crud.PostingData, 0)
		newGradList := make([]crud.PostingData, 0)

		for _, posting := range postings {
			if posting.IsIntern {
				internList = append(internList, posting)
			} else {
				newGradList = append(newGradList, posting)
			}
		}

		res, err := json.Marshal(PostResponse{internList, newGradList})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

	postPostingsHandler := func(w http.ResponseWriter, r *http.Request) {
		var postingData []crud.PostingData

		err := json.NewDecoder(r.Body).Decode(&postingData)
		if err != nil {
			fmt.Printf("error here\n")
			fmt.Printf("%v\n", postingData)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		success := InsertPosting(postingData)

		// Set Response Type on Header
		w.Header().Set("Content-Type", "application/json")

		res, err := json.Marshal(Response{Success: success})

		if !success {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}

	HandleMultiplePostingRoutes([]string{"", "/"}, getPostingsHandler, "GET", false)
	HandleMultiplePostingRoutes([]string{"", "/"}, postPostingsHandler, "POST", false)
}