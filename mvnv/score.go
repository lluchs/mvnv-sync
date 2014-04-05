package mvnv

import (
	"labix.org/v2/mgo/bson"
)

type Score struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Composer string `json:"composer,omitempty"`
	Publisher string `json:"publisher,omitempty"`
	Tags []string `json:"tags"`
	Year int `json:"year,omitempty"`
}

// Returns a slice of all scores in the database.
func GetScores() (scores []Score, err error) {
	collection := db().C("scores")
	err = collection.Find(bson.M{}).All(&scores)
	return
}
