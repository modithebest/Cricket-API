// controllers/integration_test.go
package controllers

import (
	"net/http"
	"net/http/httptest"

	"testing"

	"github.com/gin-gonic/gin"
)

func TestPost(t *testing.T) {
	// Create a new gin engine

	r := gin.Default()

	// Define the route we want to test
	r.POST("/post", Post)

	// Create a request to pass to our handler
	req, _ := http.NewRequest("POST", "/post", nil)

	// Create a ResponseRecorder to record the response
	w := httptest.NewRecorder()

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response
	if w.Code == http.StatusOK {
		t.Fatalf("Expected status code 500, got %v", w.Code)
	}

}
