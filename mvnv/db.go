package mvnv

import (
	"labix.org/v2/mgo"
)

var (
	database *mgo.Database
	session *mgo.Session
)

// Returns a connection to the mvnv database, opening it on first call.
func db() *mgo.Database {
	if database == nil {
		var err error
		session, err = mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
		database = session.DB("mvnv")
	}
	return database
}

// Closes the database session if there is one.
func CloseDB() {
	if session != nil {
		session.Close()
		session = nil
		database = nil
	}
}
