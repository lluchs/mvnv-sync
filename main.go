package main

import (
	"fmt"
	"os"
	"encoding/json"

	"mvnv-sync/mvnv"
)

func main() {
	defer mvnv.CloseDB()

	scores, err := mvnv.GetScores()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%d scores found.\n", len(scores))

	jsonString, err := json.Marshal(scores)
	if err != nil {
		fmt.Println("Marshal error: ", err)
	}
	os.Stdout.Write(jsonString)
}
