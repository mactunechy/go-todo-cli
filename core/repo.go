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

type Todo struct {
	id          int
	description string
	createdAt   string
	status      string
}

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
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: "#"},
			{Align: simpletable.AlignLeft, Text: "DESCRIPTION"},
			{Align: simpletable.AlignRight, Text: "CREATED AT"},
			{Align: simpletable.AlignRight, Text: "STATUS"},
		},
	}

	records := findAll()

	for _, record := range records {
		timeAgo := formatTime(record.createdAt)

		r := []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: fmt.Sprintf("%d", record.id)},
			{Align: simpletable.AlignLeft, Text: record.description},
			{Align: simpletable.AlignRight, Text: timeAgo},
			{Align: simpletable.AlignRight, Text: record.status},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleCompactLite)
	fmt.Println(table.String())
}

func Update(todoID int, newStatus string) error {
	dbfile, err := dbFile()
	if err != nil {
		log.Fatalln("Db not found", err)
	}
	file, err := os.OpenFile(dbfile, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("failed to open file", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Failed to read", err)
	}

	if len(rows) < todoID {
		log.Fatalf("Invalid ID. id out of range.%v", todoID)
	}

	rows[todoID-1][2] = newStatus

	file.Seek(0, 0)
	writer := csv.NewWriter(file)
	err = writer.WriteAll(rows)
	if err != nil {
		log.Fatalln("Error updating TODO -> writing", err)
	}

	return nil
}

// Private functions

func findAll() []Todo {
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

	var todos []Todo

	for idx, record := range records {
		todos = append(todos, Todo{
			id:          idx + 1,
			description: record[0],
			createdAt:   record[1],
			status:      record[2],
		})
	}
	return todos
}

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
