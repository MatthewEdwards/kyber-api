package app

import (
	"encoding/json"
	"kyber-api/datastore"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

func (app *application) handleGetAricles() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Returning all articles")

		articles := app.db.GetArticles()
		respond(w, articles, 200)
	}
}

func (app *application) handlePostAricle() http.HandlerFunc {

	var validate *validator.Validate
	validate = validator.New()

	return func(w http.ResponseWriter, r *http.Request) {
		log.Info("Adding a new article")

		var art datastore.Article

		err := json.NewDecoder(r.Body).Decode(&art)

		if err != nil {
			log.Error("Error posting article: ", err)
			respond(w, &response{Error: true, Code: http.StatusBadRequest, Message: err.Error()}, http.StatusBadRequest)
			return
		}

		t := time.Now()
		timestamp := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, time.Local).Format("2006-01-02 15:04:05")

		art.Date = timestamp

		err = validate.Struct(art)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			log.Error("Unable to validate article", validationErrors)
			respond(w, &response{Error: true, Code: http.StatusBadRequest, Message: validationErrors.Error()}, http.StatusBadRequest)
			return
		}

		err = app.db.AddArticle(art)

		if err != nil {
			log.Error("Error adding article to database: ", err)
			respond(w, &response{Error: true, Code: http.StatusBadRequest, Message: err.Error()}, http.StatusBadRequest)
			return
		}

		respond(w, &response{Error: false, Code: http.StatusOK, Message: "Article has been posted"}, http.StatusOK)
	}
}
