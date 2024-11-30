package core

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Save(todo string) error {
	dbfile, err := dbFile()
	if err != nil {
		log.Fatalln("Db not found", err)
	}

	file, err := os.OpenFile(dbfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	currentTime := time.Now().Format(time.RFC822)
	record := []string{todo, currentTime, "TODO"}

	return writer.Write(record)
}

func dbFile() (string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: %w", err)
	}

	return filepath.Join(rootDir, ".data.csv"), nil
}
