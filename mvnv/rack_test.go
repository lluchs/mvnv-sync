package mvnv

import "testing"

func TestGetRacks(t *testing.T) {
	racks, err := GetRacks()
	defer CloseDB()
	if err != nil {
		t.Fatal(err)
	}
	if len(racks) == 0 {
		t.Error("No racks returned.")
	}
}
