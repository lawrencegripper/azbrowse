package storage

import (
	"fmt"
	"log"
	"os/user"
	"time"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

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

	waitingMessageTimer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("AzBrowse is waiting for access to '~/.azbrowse.db', do you have another instance of azbrowse open?")
	})

	dbCreate, err := bolt.Open(dbLocation, 0600, nil)
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
