package main

import (
	"fmt"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"mvnv-sync/mvnv"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
			panic(err)
	}
	defer session.Close()

	scores := session.DB("mvnv").C("scores")

	score := mvnv.Score{}
	err = scores.Find(bson.M{"id": 124}).One(&score)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Score: %v\n", score)
}
