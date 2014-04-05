package mvnv

import "testing"

func TestGetBackup(t *testing.T) {
	backup, err := GetBackup()
	if err != nil {
		t.Fatal("Error: ", err)
	}
	if len(backup.Scores) == 0 {
		t.Error("Scores not properly populated.")
	}
	if len(backup.Racks) == 0 {
		t.Error("Racks not properly populated.")
	}
}
