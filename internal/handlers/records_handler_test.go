package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpereira/package_search/internal/types"
)

func TestRecordsHandler(t *testing.T) {
	// Setup
	db := setupTestDB(t)
	defer db.Close()

	err := db.Ping()
	assert.NoError(t, err)

	handler := RecordsHandler(db)

	// Test case 1: Records found
	req, err := http.NewRequest("GET", "/records?name=package2&name=package1", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var records []types.Record
	err = json.Unmarshal(rr.Body.Bytes(), &records)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(records))

	assert.Equal(t, "package1", records[0].Name)
	assert.Equal(t, "1.0.0", records[0].Version)
	assert.Equal(t, "1", records[0].Release)
	assert.Equal(t, "/usr/local", records[0].Location)
	assert.Equal(t, "Package 1 summary", records[0].Summary)

	assert.Equal(t, "package2", records[1].Name)
	assert.Equal(t, "1.1.0", records[1].Version)
	assert.Equal(t, "2", records[1].Release)
	assert.Equal(t, "/usr/local", records[1].Location)
	assert.Equal(t, "Package 2 summary", records[1].Summary)

	// Test case 2: Records not found
	req, err = http.NewRequest("GET", "/records?name=nonexistent", nil)
	assert.NoError(t, err)

	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var emptyRecords []types.Record
	err = json.Unmarshal(rr.Body.Bytes(), &emptyRecords)
	assert.NoError(t, err)

	assert.Equal(t, 0, len(emptyRecords))
}
