package models

import("gopkg.in/mgo.v2/bson")

type MongoArticle struct{
	ID     bson.ObjectId    `bson:"_id"`
	Title  string 			`bson:"title"`
	Date   string 			`bson:"date"`
	Site   string			`bson:"site"`
	URL    string 			`bson:"url"`
}