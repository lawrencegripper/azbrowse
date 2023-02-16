package storage

import (
	"fmt"
	"os/user"
	"strconv"
	"time"

	"github.com/peterbourgon/diskv"
	mockableClock "github.com/stephanos/clock"
)

var diskstore *diskv.Diskv
var clock mockableClock.Clock

const ttlLastUpdatedKey = "LastUpdated"

// GetStorageDir returns the directory that azb uses for storage
func GetStorageDir() string {
	diskLocation := "/root/.azbrowse/"
	user, err := user.Current()
	if err == nil {
		diskLocation = user.HomeDir + "/.azbrowse/"
	}
	return diskLocation
}

// LoadDB initializes and loads the DB instance
func LoadDB() {
	initDb(GetStorageDir(), mockableClock.New())
}

func initDb(location string, inputClock mockableClock.Clock) {
	clock = inputClock
	flatTransform := func(s string) []string { return []string{} }
	diskstore = diskv.New(diskv.Options{
		BasePath:     location,
		Transform:    flatTransform,
		CacheSizeMax: 1024 * 1024,
	})
}

// DeleteCache removes an item from cache
func DeleteCache(key string) error {
	if !diskstore.Has(key) {
		return nil
	}
	return diskstore.Erase(key)
}

// PutCache puts an item in the cache bucket
func PutCache(key, value string) error {
	err := diskstore.Write(key, []byte(value))
	if err != nil {
		return err
	}

	return nil
}

// GetCache gets an item from the cache bucket
func GetCache(key string) (string, error) {
	// Honor legacy - todo: fixup and return notfound error
	if !diskstore.Has(key) {
		return "", nil
	}
	result, err := diskstore.Read(key)

	if err != nil {
		return "", err
	}
	return string(result), nil
}

// GetCacheWithTTL gets an item from the cache if it's with the TTL duration. To simplify the TTL is provided by the caller
// the data store just tracks the key and it's last updated value
func GetCacheWithTTL(key string, ttl time.Duration) (valid bool, value string, err error) {
	cacheItem, err := GetCache(key)
	if err != nil {
		return false, "", err
	}

	// Empty string isn't a valid cache item
	if cacheItem == "" {
		return false, "", err
	}

	// Get the Last updated time for this cache key
	cacheItemLastUpdated, err := GetCache(key + ttlLastUpdatedKey)
	if err != nil || cacheItemLastUpdated == "" {
		return false, cacheItem, err
	}
	lastUpdatedEpoc, err := strconv.ParseInt(cacheItemLastUpdated, 10, 64)
	if err != nil {
		return false, cacheItem, fmt.Errorf("Failed to parse %v: %w", cacheItemLastUpdated, err)
	}

	// Check if the ttl has expired
	ttlExpiresAfter := time.Unix(lastUpdatedEpoc, 0).Add(ttl)
	if clock.Now().After(ttlExpiresAfter) {
		// It has
		return false, cacheItem, nil
	}

	return true, cacheItem, nil
}

// PutCacheItemForTTL puts an item in the cache bucket.
// *Warning* the current setup DOES NOT cleanup items after their TTL it only provides `GetCacheWithTTL`
// which allows the user to get the key and highlights if it's past the TTL. TTL is defined by the caller when using `GetCacheWithTTL`.
// This was used as currently no keys are noisy and require cleanup, future uses could update this to do cleanup.
func PutCacheForTTL(key, value string) error {
	// Save the Item
	err := PutCache(key, value)
	// Track when it was saved
	errUpdate := PutCache(key+ttlLastUpdatedKey, epocToString())
	if errUpdate != nil {
		return errUpdate
	}
	return err
}

func epocToString() string {
	return strconv.FormatInt(clock.Now().Unix(), 10)
}
