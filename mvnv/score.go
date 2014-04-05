package mvnv

import (
	"labix.org/v2/mgo/bson"
)

type Score struct {
	Id int
	Title string
	Composer string
	Publisher string
	Tags []string
	Year int
}

// Returns a slice of all scores in the database.
func GetScores() (scores []Score, err error) {
	collection := db().C("scores")
	err = collection.Find(bson.M{}).All(&scores)
	return
}
