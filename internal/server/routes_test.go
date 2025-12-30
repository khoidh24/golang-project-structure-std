package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// mockDB implements database.Service for testing purposes
type mockDB struct{}

func (m *mockDB) Health() map[string]string {
	return map[string]string{
		"status":  "up",
		"message": "It's healthy",
	}
}

func (m *mockDB) Close() error {
	return nil
}

func TestRoutes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Inject mock database
	s := &Server{
		db: &mockDB{},
	}

	// Register routes to get the handler
	handler := s.RegisterRoutes()

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		expectedBody   map[string]string
	}{
		{
			name:           "V1 Home Handler",
			method:         "GET",
			path:           "/v1",
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]string{"message": "Hello World"},
		},
		{
			name:           "V1 Health Handler",
			method:         "GET",
			path:           "/v1/health",
			expectedStatus: http.StatusOK,
			expectedBody:   map[string]string{"status": "up", "message": "It's healthy"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			// Validate response body JSON
			var actualBody map[string]string
			if err := json.Unmarshal(rr.Body.Bytes(), &actualBody); err != nil {
				t.Errorf("Failed to parse response body: %v", err)
			}

			for k, v := range tt.expectedBody {
				if actualBody[k] != v {
					t.Errorf("Handler returned unexpected body content for key %s: got %s want %s", k, actualBody[k], v)
				}
			}
		})
	}
}
