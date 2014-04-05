package main

import (
	"fmt"
	"os"
	"encoding/json"

	"mvnv-sync/mvnv"
)

func usage() {
	fmt.Fprintln(os.Stderr, `Usage: mvnv-sync <cmd>

Possible commands:
   backup    Print a full database backup to stdout
`)
}

func main() {
	defer mvnv.CloseDB()

	if len(os.Args) != 2 {
		usage()
		return
	}

	switch cmd := os.Args[1]; cmd {
	case "backup":
		backup, err := mvnv.GetBackup()
		if err != nil {
			panic(err)
		}

		fmt.Fprintf(os.Stderr, "%d scores and %d racks found.\n", len(backup.Scores), len(backup.Racks))

		jsonString, err := json.Marshal(backup)
		if err != nil {
			fmt.Println("Marshal error: ", err)
		}
		os.Stdout.Write(jsonString)
	default:
		fmt.Fprintln(os.Stderr, "Invalid command ", cmd)
		usage()
	}
}
