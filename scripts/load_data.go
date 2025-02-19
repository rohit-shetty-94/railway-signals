package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pushpanthraj/crosstech/railway-signals/config"
	"github.com/pushpanthraj/crosstech/railway-signals/entity/track"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	file, err := os.Open("../data.json")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)
	var tracks []track.Track

	if err := json.Unmarshal(byteValue, &tracks); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	for _, track := range tracks {
		_, err := db.Model(&track).Insert()
		if err != nil {
			log.Fatalf("Failed to insert track: %v", err)
		}

		for _, signal := range track.Signals {
			signal.TrackID = track.ID
			_, err := db.Model(&signal).Insert()
			if err != nil {
				log.Printf("Failed to insert signal: %v", err)
			}
		}
	}

	fmt.Println("Data inserted successfully!")
}
