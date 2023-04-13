package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecordHandler(t *testing.T) {
	// Setup
	db := setupTestDB(t)
	defer db.Close()

	err := db.Ping()
	assert.NoError(t, err)

	handler := recordHandler(db)

	// Test case 1: Record found
	req, err := http.NewRequest("GET", "/record?name=package1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var record Record
	err = json.Unmarshal(rr.Body.Bytes(), &record)
	assert.NoError(t, err)

	assert.Equal(t, "package1", record.Name)
	assert.Equal(t, "1.0.0", record.Version)
	assert.Equal(t, "1", record.Release)
	assert.Equal(t, "/usr/local", record.Location)
	assert.Equal(t, "Package 1 summary", record.Summary)

	// Test case 2: Record not found
	req, err = http.NewRequest("GET", "/record?name=nonexistent", nil)
	assert.NoError(t, err)

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}
