package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/pushpanthraj/crosstech/railway-signals/entity/signal"
	"github.com/pushpanthraj/crosstech/railway-signals/entity/track"
)

func GetAllSignals(db *pg.DB) ([]signal.Signal, error) {
	var signals []signal.Signal
	err := db.Model(&signals).Select()
	return signals, err
}

func GetTracksBySignal(db *pg.DB, signalID int) ([]track.Track, error) {
	var tracks []track.Track
	err := db.Model(&tracks).
		Join("JOIN signals ON signals.track_id = tracks.id").
		Where("signals.signal_id = ?", signalID).
		Select()

	return tracks, err
}

func GetTracksBySignal1(db *pg.DB, signalID int) ([]track.Track, error) {
	var tracks []track.Track
	err := db.Model(&tracks).
		Join("JOIN signals ON signals.track_id = tracks.id").
		Where("signals.signal_id = ?", signalID).
		Select()
	return tracks, err
}

func AssociateTrackWithSignal(db *pg.DB, trackID, signalID int) error {
	_, err := db.Exec("UPDATE signals SET track_id = ? WHERE id = ?", trackID, signalID)
	return err
}

func RemoveTrackSignalAssociation(db *pg.DB, trackID, signalID int) error {
	_, err := db.Exec("UPDATE signals SET track_id = NULL WHERE id = ? AND track_id = ?", signalID, trackID)
	return err
}
