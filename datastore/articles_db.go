package datastore

import (
	"kyber-api/models"
	"kyber-api/utils"

	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

// GetArticles will return a list of all clients in the database
func (c *MgoConnection) GetArticles() (articles []models.MongoArticle) {
	log.Info("[Datastore] GetArticles")
	session, articlesCollection, err := c.getSession()

	defer session.Close()

	if err != nil {
		log.Error("[Datastore] GetArticles Error: ", err)
		return nil
	}

	// Get the articles and sort them in reverse order
	dberr := articlesCollection.Find(nil).Sort("-_id").All(&articles)

	if dberr != nil {
		log.Error("[Datastore] GetArticles Error: ", dberr)
		return nil
	}

	return articles
}

// AddArticle adds a new article to the database
func (c *MgoConnection) AddArticle(article models.Article) (err error) {
	log.Info("[Datastore] AddArticle")
	session, articlesCollection, err := c.getSession()

	defer session.Close()

	if err != nil {
		log.Info("[Datastore] AddArticle Error: ", err)
		return err
	}

	// Convert the Article into a MongoArticle and add it to the database
	err = articlesCollection.Insert(
		&models.MongoArticle{
			ID:    bson.NewObjectId(),
			Title: article.Title,
			Date:  utils.GetTimestamp(),
			Site:  article.Site,
			URL:   article.URL,
		},
	)

	if err != nil {
		log.Info("[Datastore] AddArticle Error: ", err)
		return err
	}

	return nil
}
