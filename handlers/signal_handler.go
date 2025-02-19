package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"

	"github.com/pushpanthraj/crosstech/railway-signals/constants"
	"github.com/pushpanthraj/crosstech/railway-signals/repository"
	"github.com/pushpanthraj/crosstech/railway-signals/utils"
)

// Get all signals
func GetSignals(db *pg.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		signals, err := repository.GetAllSignals(db)
		if err != nil {
			return utils.SendResponse(c, http.StatusInternalServerError, constants.ErrRecordsRetrivedFail, err.Error())
		}
		return utils.SendResponse(c, http.StatusOK, constants.MsgsReadSuccess, signals)
	}
}

// Get all tracks associated with a signal
func GetTracksBySignal(db *pg.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, constants.ErrInvalidInput)
		}

		tracks, err := repository.GetTracksBySignal(db, id)
		if err != nil {
			return utils.SendResponse(c, http.StatusNotFound, constants.ErrRecordsRetrivedFail, err.Error())
		}

		return utils.SendResponse(c, http.StatusOK, constants.MsgsReadSuccess, tracks)
	}
}

// Associate a track with a signal
func AssociateTrackWithSignal(db *pg.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		trackID, _ := strconv.Atoi(c.Param("track_id"))
		signalID, _ := strconv.Atoi(c.Param("signal_id"))

		err := repository.AssociateTrackWithSignal(db, trackID, signalID)
		if err != nil {
			return utils.SendResponse(c, http.StatusInternalServerError, constants.ErrRecordsRetrivedFail, err.Error())
		}

		// here retrive signal back in response
		return utils.SendResponse(c, http.StatusOK, constants.TrackSignalAssociatedSuccess, nil)
	}
}

// Remove the association between a signal and a track
func RemoveTrackSignalAssociation(db *pg.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		trackID, _ := strconv.Atoi(c.Param("track_id"))
		signalID, _ := strconv.Atoi(c.Param("signal_id"))

		err := repository.RemoveTrackSignalAssociation(db, trackID, signalID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		// here retrive signal back in response
		return utils.SendResponse(c, http.StatusOK, constants.TrackSignalAssociatedRemoved, nil)
	}
}
