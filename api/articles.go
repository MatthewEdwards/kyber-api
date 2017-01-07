package api

import (
	"encoding/json"
	"net/http"
	"kyber-api/models"
	log "github.com/Sirupsen/logrus"
)

// GetArticlesHandler retrives a list of all the articles from the database
func (DS *DStore)GetArticlesHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("[API] GetArticlesHandler ", r.URL)
	w.Header().Set("Content-Type", "application/json")

	var articles []models.MongoArticle

	// Get the articles and return them as JSON
	articles = DS.connection.GetArticles()
	j, _ := json.Marshal(articles)
	
	w.Header().Set("Access-Control-Allow-Origin", "*")	
	w.Write(j)
}

// PostArticlesHandler adds a new article to the database
func (DS *DStore)PostArticlesHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var a models.Article

	// Read the post data and store it a new article object
	decoder := json.NewDecoder(r.Body)
	jsonError := decoder.Decode(&a)

	// Return a StatusBadRequest if there was an error
	if jsonError != nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error bad request"))
		log.Error("[PostArticlesHandler] Error: ", jsonError)
		return
	}
	 	
	defer r.Body.Close()
	
	
	// Add the article to db
	dbError := DS.connection.AddArticle(a)

	// Return a StatusBadRequest if there was an error
	if dbError != nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error bad request"))
		log.Error("[PostArticlesHandler] Error: ", dbError)
		return
	}

	// Return a 201 success if there was no error
	j, _ := json.Marshal(a)
	w.WriteHeader(201)
	w.Write(j)
}
