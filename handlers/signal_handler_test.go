package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pushpanthraj/crosstech/railway-signals/utils"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Mock Database Connection
func mockDB() *pg.DB {
	return &pg.DB{}
}

// Test GetSignals Handler
func TestGetSignals(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/signals", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := GetSignals(mockDB())

	// Invoke Handler
	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response utils.APIResponse
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, "Signals retrieved successfully", response.Header.Message)
	}
}

// Test GetTracksBySignal Handler
func TestGetTracksBySignal(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/signals/1/tracks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	handler := GetTracksBySignal(mockDB())

	// Invoke Handler
	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response utils.APIResponse
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, "Tracks retrieved successfully", response.Header.Message)
	}
}

// Test AssociateTrackWithSignal Handler
func TestAssociateTrackWithSignal(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/tracks/1/signals/2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("track_id", "signal_id")
	c.SetParamValues("1", "2")

	handler := AssociateTrackWithSignal(mockDB())

	// Invoke Handler
	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response utils.APIResponse
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, "Track associated with signal successfully", response.Header.Message)
	}
}

// Test RemoveTrackSignalAssociation Handler
func TestRemoveTrackSignalAssociation(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/tracks/1/signals/2", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("track_id", "signal_id")
	c.SetParamValues("1", "2")

	handler := RemoveTrackSignalAssociation(mockDB())

	// Invoke Handler
	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response utils.APIResponse
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, "Track-signal association removed successfully", response.Header.Message)
	}
}
