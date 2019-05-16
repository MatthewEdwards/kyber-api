package app

import (
	"encoding/json"
	"kyber-api/datastore"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (app *application) handleGetAricles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Returning all articles")

		articles := app.db.GetArticles()
		response(w, articles, 200)
	}
}

func (app *application) handlePostAricle() http.HandlerFunc {

	type error struct {
		Error   bool   `json:"error"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	type success struct {
		Error   bool   `json:"error"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Posting a new article")

		var art datastore.Article

		err := json.NewDecoder(r.Body).Decode(&art)

		if err != nil {
			log.Error("Error posting article: ", err)
			response(w, &error{Error: true, Code: http.StatusBadRequest, Message: err.Error()}, http.StatusBadRequest)
			return
		}

		err = app.db.AddArticle(art)

		if err != nil {
			log.Error("Error adding article to database: ", err)
			response(w, &error{Error: true, Code: http.StatusBadRequest, Message: err.Error()}, http.StatusBadRequest)
			return
		}

		response(w, &success{Error: false, Code: http.StatusOK, Message: "Article has been posted"}, http.StatusOK)
	}
}