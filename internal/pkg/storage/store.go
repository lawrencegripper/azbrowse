package storage

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os/user"
	"time"
)

var db *bolt.DB

// LoadDB initializes and loads the DB instance
func LoadDB() {
	fmt.Println("Loading db ...")
	dbLocation := "/root/.azbrowse.db"
	user, err := user.Current()
	if err == nil {
		dbLocation = user.HomeDir + "/.azbrowse.db"
	}

	suppressWaitingMessage := false
	go func() {
		time.Sleep(2 * time.Second)
		if !suppressWaitingMessage {
			fmt.Println("AzBrowse is waiting for access to '~/.azbrowse.db', do you have another instance of azbrowse open?")
		}
	}()

	dbCreate, err := bolt.Open(dbLocation, 0600, nil)
	suppressWaitingMessage = true
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

// ClearResources removes all resources from the
func ClearResources() error {
	return db.Update(func(tx *bolt.Tx) error {
		return tx.DeleteBucket([]byte("search"))
	})
}

// PutResourceBatch puts a batch of resources into the boltdb
func PutResourceBatch(key string, value []Resource) error {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("search"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		return err
	}

	return db.Batch(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("search"))
		bytes, err := json.Marshal(value)
		if err != nil {
			return err
		}
		err = b.Put([]byte(key), bytes)
		if err != nil {
			return err
		}
		return nil
	})
}

// GetAllResources returns all the resources seen in the last crawl
func GetAllResources() ([]Resource, error) {
	resources := []Resource{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("search"))
		return b.ForEach(func(key, value []byte) error {
			var storedResources []Resource
			err := json.Unmarshal([]byte(value), &storedResources)
			if err != nil {
				return err
			}

			resources = append(resources, storedResources...)
			return nil
		})
	})
	return resources, err
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
