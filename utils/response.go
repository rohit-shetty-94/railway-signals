package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/pushpanthraj/crosstech/railway-signals/entity/signal"
	"github.com/pushpanthraj/crosstech/railway-signals/entity/track"
)

// Response struct for all API responses
type APIResponse struct {
	Header struct {
		Message    string `json:"message"`
		Count      int    `json:"count"`
		StatusCode int    `json:"status_code"`
	} `json:"header"`
	Data interface{} `json:"data"`
}

// Function to send a response
func SendResponse(c echo.Context, status int, message string, data interface{}) error {
	response := APIResponse{
		Data: data,
	}
	response.Header.Message = message
	response.Header.StatusCode = status

	// // If data is a slice, count the number of items
	// if d, ok := data.([]interface{}); ok {
	// 	response.Header.Count = len(d)
	// } else {
	// 	response.Header.Count = 1 // Single object
	// }

	if data != nil {
		switch v := data.(type) {
		case []interface{}:
			response.Header.Count = len(v)
		case []track.Track:
			response.Header.Count = len(v)
		case []signal.Signal:
			response.Header.Count = len(v)
		case *track.Track:
			response.Header.Count = 1
		case *signal.Signal:
			response.Header.Count = 1
		}
	}

	return c.JSON(status, response)
}
