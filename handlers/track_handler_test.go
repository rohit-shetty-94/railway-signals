package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pushpanthraj/crosstech/railway-signals/utils"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Test GetTracks Handler
func TestGetTracks(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/tracks", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := GetTracks(mockDB())

	// Invoke Handler
	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response utils.APIResponse
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, "Tracks retrieved successfully", response.Header.Message)
	}
}

// Test GetTrackByID Handler
func TestGetTrackByID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/tracks/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	handler := GetTrackByID(mockDB())

	// Invoke Handler
	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response utils.APIResponse
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, "Track retrieved successfully", response.Header.Message)
	}
}

// Test CreateTrack Handler
func TestCreateTrack(t *testing.T) {
	e := echo.New()
	payload := `{"source": "Test Source", "target": "Test Target"}`
	req := httptest.NewRequest(http.MethodPost, "/tracks", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := CreateTrack(mockDB())

	// Invoke Handler
	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		var response utils.APIResponse
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, "Track created successfully", response.Header.Message)
	}
}

// Test UpdateTrack Handler
func TestUpdateTrack(t *testing.T) {
	e := echo.New()
	payload := `{"source": "Updated Source", "target": "Updated Target"}`
	req := httptest.NewRequest(http.MethodPut, "/tracks/1", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	handler := UpdateTrack(mockDB())

	// Invoke Handler
	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response utils.APIResponse
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, "Track updated successfully", response.Header.Message)
	}
}

// Test DeleteTrack Handler
func TestDeleteTrack(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/tracks/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	handler := DeleteTrack(mockDB())

	// Invoke Handler
	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var response utils.APIResponse
		json.Unmarshal(rec.Body.Bytes(), &response)

		assert.Equal(t, "Track deleted successfully", response.Header.Message)
	}
}
