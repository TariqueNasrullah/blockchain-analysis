package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var files []string
	err := filepath.Walk("gen_time/", func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	csvFile, err := os.Create("dataset.csv")
	if err != nil {
		panic(err)
	}

	w := csv.NewWriter(csvFile)

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Printf("Error opening file: %v\n", err)
			continue
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			data := strings.Split(scanner.Text(), ",")
			w.Write(data)

		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
