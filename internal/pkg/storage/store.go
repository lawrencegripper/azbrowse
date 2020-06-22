package storage

import (
	"fmt"
	"log"
	"os/user"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	mockableClock "github.com/stephanos/clock"
)

var db *bolt.DB
var clock mockableClock.Clock

const ttlLastUpdatedKey = "LastUpdated"

// CloseDB closes the db
func CloseDB() {
	err := db.Close()
	if err != nil {
		panic(err)
	}
}

// LoadDB initializes and loads the DB instance
func LoadDB() {
	fmt.Println("Loading db ...")
	dbLocation := "/root/.azbrowse.db"
	user, err := user.Current()
	if err == nil {
		dbLocation = user.HomeDir + "/.azbrowse.db"
	}

	initDb(dbLocation, mockableClock.New())
}

func initDb(location string, inputClock mockableClock.Clock) {
	clock = inputClock
	waitingMessageTimer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("AzBrowse is waiting for access to '~/.azbrowse.db', do you have another instance of azbrowse open?")
	})
	dbCreate, err := bolt.Open(location, 0600, nil)
	waitingMessageTimer.Stop()
	if err != nil {
		log.Fatal(err)
	}

	db = dbCreate

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("cache"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		_, err = tx.CreateBucketIfNotExists([]byte("search"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Loading db complete")
}

// PutCache puts an item in the cache bucket
func PutCache(key, value string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("cache"))
		err := b.Put([]byte(key), []byte(value))
		return err
	})
}

// GetCache gets an item from the cache bucket
func GetCache(key string) (string, error) {
	var s []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("cache"))
		v := b.Get([]byte(key))
		s = v
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("Failed to find item: %v", err)
	}
	return string(s), nil
}

// GetCacheIfWithinTTL gets an item from the cache if it's with the TTL duration. To simplify the TTL is provided by the caller
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
		return false, cacheItem, errors.Wrapf(err, "Failed to parse %v", cacheItemLastUpdated)
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
