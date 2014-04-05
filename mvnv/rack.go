package mvnv

import (
	"labix.org/v2/mgo/bson"
)

type Rack struct {
	Name string `json:"name"`
	Contents struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"contents"`
}

// Returns a slice of all racks in the database.
func GetRacks() (racks []Rack, err error) {
	collection := db().C("racks")
	err = collection.Find(bson.M{}).All(&racks)
	return
}
