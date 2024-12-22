package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"calc_service/internal/handler"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		method     string
		body       interface{}
		statusCode int
		response   map[string]interface{}
	}{
		{
			method:     http.MethodGet,
			body:       nil,
			statusCode: http.StatusMethodNotAllowed,
			response:   map[string]interface{}{"error": "Method not allowed"},
		},
		{
			method:     http.MethodPost,
			body:       nil,
			statusCode: http.StatusBadRequest,
			response:   map[string]interface{}{"error": "Invalid request format"},
		},
		{
			method:     http.MethodPost,
			body:       map[string]string{"expression": ""},
			statusCode: http.StatusUnprocessableEntity,
			response:   map[string]interface{}{"error": "Expression is not valid: empty expression"},
		},
		{
			method:     http.MethodPost,
			body:       map[string]string{"expression": "2 + 2"},
			statusCode: http.StatusOK,
			response:   map[string]interface{}{"result": 4.0},
		},
		{
			method:     http.MethodPost,
			body:       map[string]string{"expression": "10 / 0"},
			statusCode: http.StatusUnprocessableEntity,
			response:   map[string]interface{}{"error": "Expression is not valid: division by zero"},
		},
	}

	for _, test := range tests {
		t.Run(test.method, func(t *testing.T) {
			var body []byte
			if test.body != nil {
				body, _ = json.Marshal(test.body)
			}

			req := httptest.NewRequest(test.method, "/calculate", bytes.NewReader(body))
			res := httptest.NewRecorder()

			handler.CalculateHandler(res, req)

			if res.Code != test.statusCode {
				t.Errorf("unexpected status code: got %v, want %v", res.Code, test.statusCode)
			}

			var response map[string]interface{}
			_ = json.NewDecoder(res.Body).Decode(&response)

			for key, value := range test.response {
				if response[key] != value {
					t.Errorf("unexpected response %q: got %v, want %v", key, response[key], value)
				}
			}
		})
	}
}
