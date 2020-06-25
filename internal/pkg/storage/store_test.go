package storage

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	mockableClock "github.com/stephanos/clock"
)

var mockClock = mockableClock.NewMock()
var testTime = time.Date(2019, 01, 01, 01, 00, 00, 00, time.UTC)

func TestCacheWithTTL(t *testing.T) {
	// Create a test instance of the DB
	dirName, err := ioutil.TempDir(os.TempDir(), "azb-storagetests.db")
	if err != nil {
		log.Fatal(err)
	}
	// defer os.(dirName) //nolint: errcheck
	initDb(dirName, mockClock)

	tests := []struct {
		// Name is used to identify the test, as the cache key and as the cache value
		// it should be unique in the test set.
		name           string
		ttl            time.Duration
		fastForward    time.Duration
		putValue       bool
		wantValid      bool
		wantValueMatch bool
		wantErr        bool
	}{
		{
			name:           "InsertRetrieveExpectValid",
			ttl:            time.Hour,
			fastForward:    time.Minute * 45,
			wantValid:      true,
			wantValueMatch: true,
			wantErr:        false,
		},
		{
			name:           "InsertRetrieveExpectInvalid",
			ttl:            time.Hour,
			fastForward:    time.Hour * 2,
			wantValid:      false,
			wantValueMatch: true,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := PutCacheForTTL(tt.name, tt.name)
			if err != nil {
				t.Errorf("PutCacheForTTL errored = %v, want no error", err)
				return
			}
			mockClock.Set(testTime.Add(tt.fastForward))

			gotValid, gotValue, err := GetCacheWithTTL(tt.name, tt.ttl)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetCacheWithTTL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotValid != tt.wantValid {
				t.Errorf("GetCacheWithTTL() gotValid = %v, want %v", gotValid, tt.wantValid)
			}

			// Note to avoid duplication the tt.name field is set as the value of the cache item too
			if gotValue != tt.name && tt.wantValueMatch {
				t.Errorf("GetCacheWithTTL() gotValue = %v, want %v", gotValue, tt.name)
			}
		})
	}
}

func TestGetCacheWithTTL_withNonexistentKey_ExpectErr(t *testing.T) {
	// Create a test instance of the DB
	file, err := ioutil.TempFile(os.TempDir(), "azb-storagetests.db")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name()) //nolint: errcheck
	initDb(file.Name(), mockClock)

	isValid, _, err := GetCacheWithTTL("keydoesntexist", time.Hour)

	if isValid {
		t.Error("Expect invalid cache result. Got valid")
	}

	if err != nil {
		t.Errorf("Expected no error but got err = %+v", err)
	}
}
