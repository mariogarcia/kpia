package api

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func applyMigraton(db *sql.DB) func(path string, file os.FileInfo, err error) error {
	return func(path string, file os.FileInfo, err error) error {
		if file.IsDir() {
			return nil
		}

		fileName := file.Name()
		fileBytes, fileError := ioutil.ReadFile(path)

		if fileError != nil {
			log.Fatalf("migrations: error reading file %s\n", fileName)
		}

		log.Printf("migrations: applying %s\n", file.Name())
		tx, err := db.Begin()
		if err == nil {
			_, execError := db.Exec(string(fileBytes))
			if execError != nil {
				log.Fatalf("migrations: error applying migration %s\n", fileName)
			}
		}
		tx.Commit()

		return nil
	}
}

// LoadFixtures loads fixtures entries
func LoadFixtures(db *sql.DB) {
	log.Println("migrations: init...")
	directory := "./migrations"
	files, _ := ioutil.ReadDir(directory)
	log.Printf("migrations: found %d files", len(files))

	filepath.Walk(directory, applyMigraton(db))
}
