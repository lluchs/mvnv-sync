package mvnv

type Backup struct {
	Scores []Score `json:"scores"`
	Racks []Rack `json:"racks"`
}

// Returns a fully populated backup of the database.
func GetBackup() (backup Backup, err error) {
	backup.Scores, err = GetScores()
	if err != nil {
		return
	}
	backup.Racks, err = GetRacks()
	return
}
