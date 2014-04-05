package mvnv

import "testing"

func TestGetScores(t *testing.T) {
	scores, err := GetScores()
	defer CloseDB()
	if err != nil {
		t.Fatal(err)
	}
	if len(scores) == 0 {
		t.Error("No scores returned.")
	}
}
