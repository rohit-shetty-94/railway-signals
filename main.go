package main

import (
	"github.com/pushpanthraj/crosstech/railway-signals/config"
	"github.com/pushpanthraj/crosstech/railway-signals/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	e := echo.New()

	// Track Routes
	e.GET("/tracks", handlers.GetTracks(db))          // Get all tracks with signals
	e.GET("/tracks/:id", handlers.GetTrackByID(db))   // Get a specific track by ID with signals
	e.POST("/tracks", handlers.CreateTrack(db))       // Create a new track
	e.PUT("/tracks/:id", handlers.UpdateTrack(db))    // Update an existing track
	e.DELETE("/tracks/:id", handlers.DeleteTrack(db)) // Delete a track by ID

	// Signal Routes
	e.GET("/signals", handlers.GetSignals(db))                   // Get all signals
	e.GET("/signals/:id/tracks", handlers.GetTracksBySignal(db)) // Get all tracks associated with a signal

	// Signal-Track Relationship Routes
	e.POST("/tracks/:track_id/signals/:signal_id", handlers.AssociateTrackWithSignal(db))       // Associate a track with a signal
	e.DELETE("/tracks/:track_id/signals/:signal_id", handlers.RemoveTrackSignalAssociation(db)) // Remove association between a track and a signal

	e.Logger.Fatal(e.Start(":8080"))
}
