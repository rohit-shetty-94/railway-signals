package handlers

import (
	"net/http"
	"strconv"

	"github.com/pushpanthraj/crosstech/railway-signals/constants"
	"github.com/pushpanthraj/crosstech/railway-signals/entity/track"
	"github.com/pushpanthraj/crosstech/railway-signals/repository"
	"github.com/pushpanthraj/crosstech/railway-signals/utils"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
)

// Get all tracks with signals
func GetTracks(db *pg.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		tracks, err := repository.GetAllTracks(db)
		if err != nil {
			return utils.SendResponse(c, http.StatusInternalServerError, constants.ErrRecordsRetrivedFail, err.Error())
		}
		return utils.SendResponse(c, http.StatusOK, constants.MsgsReadSuccess, tracks)
	}
}

// Get a single track by ID with its signals
func GetTrackByID(db *pg.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, constants.ErrInvalidInput)
		}

		track, err := repository.GetTrackByID(db, id)
		if err != nil {
			return utils.SendResponse(c, http.StatusNotFound, constants.ErrRecordNotFound, err.Error())
		}

		return utils.SendResponse(c, http.StatusOK, constants.MsgReadSuccess, track)
	}
}

// Create a new track
func CreateTrack(db *pg.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		track := new(track.Track)
		if err := c.Bind(track); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		// for _, val := range track.Signals {
		// 	val.TrackID = track.ID
		// }
		err := repository.CreateTrack(db, track)
		if err != nil {
			return utils.SendResponse(c, http.StatusInternalServerError, constants.ErrRecordCreateFail, err.Error())
		}

		// check if signal exist with same track_id

		return utils.SendResponse(c, http.StatusCreated, constants.MsgCreateSuccess, track)
	}
}

// Update an existing track
func UpdateTrack(db *pg.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, constants.ErrInvalidInput)
		}

		track := new(track.Track)
		if err := c.Bind(track); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		track.ID = id
		err = repository.UpdateTrack(db, track)
		if err != nil {
			return utils.SendResponse(c, http.StatusInternalServerError, constants.ErrRecordUpdateFail, err.Error())
		}

		return utils.SendResponse(c, http.StatusOK, constants.MsgUpdateSuccess, track)
	}
}

// Delete a track by ID
func DeleteTrack(db *pg.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, constants.ErrInvalidInput)
		}

		err = repository.DeleteTrack(db, id)
		if err != nil {
			return utils.SendResponse(c, http.StatusInternalServerError, constants.ErrRecordDeleteFail, err.Error())
		}

		return utils.SendResponse(c, http.StatusOK, constants.MsgDeleteSuccess, nil)
	}
}
