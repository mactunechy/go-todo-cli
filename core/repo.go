package core

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/mergestat/timediff"
)

func Save(todo string) error {
	dbfile, err := dbFile()
	if err != nil {
		log.Fatalln("Db not found", err)
	}

	file, err := os.OpenFile(dbfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("Error opening db file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	currentTime := time.Now().Format(time.RFC822)
	record := []string{todo, currentTime, "TODO"}

	return writer.Write(record)
}

func List() {
	dbfile, err := dbFile()
	if err != nil {
		log.Fatalln("Db not found", err)
	}

	file, err := os.Open(dbfile)
	if err != nil {
		log.Fatalln("Error opening db file", err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Error reading records from db file", err)
	}
	defer file.Close()

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: "#"},
			{Align: simpletable.AlignLeft, Text: "DESCRIPTION"},
			{Align: simpletable.AlignRight, Text: "CREATED AT"},
			{Align: simpletable.AlignRight, Text: "STATUS"},
		},
	}

	for idx, record := range records {
		timeAgo := formatTime(record[1])

		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", idx+1)},
			{Align: simpletable.AlignLeft, Text: record[0]},
			{Align: simpletable.AlignRight, Text: timeAgo},
			{Align: simpletable.AlignRight, Text: record[2]},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}

// Private methods

func formatTime(t string) string {
	todoTime, err := time.Parse(time.RFC822, t)
	if err != nil {
		return t
	}
	return timediff.TimeDiff(todoTime)
}

func dbFile() (string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: %w", err)
	}

	return filepath.Join(rootDir, ".data.csv"), nil
}
