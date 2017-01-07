package datastore

import (
	"time"
	"gopkg.in/mgo.v2/bson"
	"kyber-api/models"
	log "github.com/Sirupsen/logrus"
)

// getTimestamp will return the current time as a ISO 8601 string
func getTimestamp() string{
	t := time.Now()
	timestamp := time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		0,
		time.Local,
	).Format("2006-01-02 15:04:05")
	
	return timestamp
}

// GetArticles will return a list of all clients in the database
func (c *MgoConnection) GetArticles() (articles []models.MongoArticle) {
	log.Info("[Datastore] GetArticles")
	session, articlesCollection, err := c.getSession()
	
	if err != nil{
		log.Error("[Datastore] GetArticles Error: ", err)
		return nil
	}

	defer session.Close()

	// Get the articles and sort them in reverse order
	dberr := articlesCollection.Find(nil).Sort("-_id").All(&articles)
	
	if dberr != nil{
		log.Error("[Datastore] GetArticles Error: ", dberr)
		return nil
	}

	return articles
}

// AddArticle adds a new article to the database
func (c *MgoConnection) AddArticle(article models.Article) (err error) {
	log.Info("[Datastore] AddArticle")
	session, articlesCollection, err := c.getSession()
	
	if err != nil{
		log.Info("[Datastore] AddArticle Error: ", err)
		return err
	}

	defer session.Close()
	
	// Convert the Article into a MongoArticle and add it to the database
	err = articlesCollection.Insert(
		&models.MongoArticle{
			ID: 	 bson.NewObjectId(),
			Title:   article.Title,
			Date:    getTimestamp(),
			Site:    article.Site,
			URL:     article.URL,
		},
	)

	if err != nil{
		log.Info("[Datastore] AddArticle Error: ", err)
		return err
	}

	return nil
}
