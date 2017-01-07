package datastore

import (
    "gopkg.in/mgo.v2"
    log "github.com/Sirupsen/logrus"
)

type MgoConnection struct {
	session *mgo.Session
}

// NewDBConnection will return a connection to the database
func NewDBConnection()(conn *MgoConnection){
	conn = new(MgoConnection)
	conn.DBConnect()

	if conn.session == nil{
		log.Fatal("[DBConnect] Error unable to connect to the database")
	}
	return conn
}

// DBConnect will connect to the database
func (c *MgoConnection) DBConnect() (err error) {
	log.Info("[Datastore] DBConnect")

	c.session, err = mgo.Dial("mongodb://localhost")
	
	if err != nil {
		log.Error("[Datastore] DBConnect Error: ", err)
		return err
	}

	articlesCollection := c.session.DB("KyberDB").C("ArticlesCollection")

	if articlesCollection == nil{
		log.Error("[Datastore] DBConnect Error")
	}
	
	return
}

func (c *MgoConnection) getSession()(session *mgo.Session, articlesCollection *mgo.Collection, err error) {
	log.Info("[Datastore] getSession")
	
	if c.session != nil{
		session = c.session.Copy()
		articlesCollection = session.DB("KyberDB").C("ArticlesCollection")
	} else{
		log.Error("[Datastore] getSession Error")
	}
	
	return
}