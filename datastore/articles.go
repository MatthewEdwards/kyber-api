package datastore

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Article will hold details for articles that are stored in the database
type Article struct {
	_ID   primitive.ObjectID `bson:"_id"`
	Title string             `bson:"title" validate:"required"`
	Source string 			 `bson:"source" validate:"required"`
	Date string 			 `bson:"date"`
	URL   string             `bson:"url" validate:"required,url"`
}

// GetArticles will return all of the articles stored in the table
func (c *MongoDB) GetArticles() (articles []Article) {
	log.Info("Retrieving articles from the database")

	collection := c.Session.Database("Kyber").Collection("Articles")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	cur, err := collection.Find(ctx, bson.D{})

	defer cur.Close(ctx)
	defer cancel()

	if err != nil {
		log.Error("Unable to retrieving articles from the database", err)
	}

	for cur.Next(ctx) {
		var article Article
		cur.Decode(&article)
		articles = append(articles, article)
	}

	return articles
}

// AddArticle will insert an article into the database
func (c *MongoDB) AddArticle(article Article) (err error) {
	log.Info("Adding a new article to the database")

	collection := c.Session.Database("Kyber").Collection("Articles")
	_, err = collection.InsertOne(context.TODO(), article)

	if err != nil {
		log.Error("Unable to add the article to the database", err)
	}

	return nil
}
