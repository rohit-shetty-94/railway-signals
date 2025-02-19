package repository

import (
	"github.com/go-pg/pg/v10"
	"github.com/pushpanthraj/crosstech/railway-signals/entity/track"
)

func GetAllTracks(db *pg.DB) ([]track.Track, error) {
	var tracks []track.Track
	err := db.Model(&tracks).Relation("Signals").Select()
	return tracks, err
}

func GetTrackByID(db *pg.DB, id int) (*track.Track, error) {
	track := &track.Track{ID: id}
	err := db.Model(track).Relation("Signals").WherePK().Select()
	return track, err
}

func CreateTrack(db *pg.DB, track *track.Track) error {
	_, err := db.Model(track).Insert()
	return err
}

func UpdateTrack(db *pg.DB, track *track.Track) error {
	_, err := db.Model(track).WherePK().Update()
	return err
}

func DeleteTrack(db *pg.DB, id int) error {
	_, err := db.Model((*track.Track)(nil)).Where("id = ?", id).Delete()
	return err
}
