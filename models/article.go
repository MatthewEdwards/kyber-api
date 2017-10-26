package models

import(
	"gopkg.in/mgo.v2/bson"
)

// Article is used to store a news article
type Article struct {
	Title string `json:"title"`
	Site  string `json:"site"`
	URL   string `json:"url"`
}

// MongoArticle is the database schema for an article
type MongoArticle struct {
	ID    bson.ObjectId `bson:"_id"`
	Title string        `bson:"title"`
	Date  string        `bson:"date"`
	Site  string        `bson:"site"`
	URL   string        `bson:"url"`
}
