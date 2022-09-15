package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/AksAman/tri/utils"
	"github.com/boltdb/bolt"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	utils.InitializeLogger("tri-db.log")
	logger = utils.Logger
}

type ByteToStringConverter func([]byte) string

func main() {

	dbName := "test.db"
	utils.Title("Opening database " + dbName)
	db, err := bolt.Open(dbName, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		logger.Errorf("Error while opening db at %q: %v", dbName, err)
	}
	defer func() {
		utils.Title("Closing database " + dbName)
		db.Close()
		logger.Debugf("%q DB Closed", dbName)
	}()

	logger.Debugf("DB opened %v", db)

	bucketName := "test"

	// creating a test bucket
	err = createBucket(bucketName, db)
	if err != nil {
		logger.Errorf("error while creating bucket(%q): %v", bucketName, err)
	}

	// add some key-value pairs
	values := map[string]int{
		"ans1": 50,
		"ans2": 100,
	}
	err = addMapToBucket(bucketName, db, values)
	if err != nil {
		logger.Errorf("error while adding data to bucket(%q): %v, rolling back", bucketName, err)
	}

	// try reading data
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	keys = append(keys, "invalidkey") // this key does not exist

	err = readDataFromDB(bucketName, db, keys)
	if err != nil {
		logger.Errorf("error while reading data from bucket(%q): %v", bucketName, err)
	}

	// iterate over all key:value pairs and show
	err = showAllData(bucketName, db, nil, func(b []byte) string {
		return strconv.Itoa(utils.BytesToInt(b))
	})
	if err != nil {
		logger.Errorf("error while showing data from bucket(%q): %v", bucketName, err)
	}
	runUserExample(db)
}

func createBucket(bucketName string, db *bolt.DB) error {
	utils.Title("Creating bucket")
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		logger.Debugf("bucket created: %v\n", bucketName)
		return nil
	})
	return err
}

func addMapToBucket(bucketName string, db *bolt.DB, data map[string]int) error {
	utils.Title("Adding data to bucket")
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("bucket:%q does not exists", bucketName)
		}
		for k, v := range data {
			err := b.Put([]byte(k), utils.IntToBytes(v))
			if err != nil {
				return err
			}
			logger.Debugf("Added %q:%d", k, v)
		}
		return nil
	})
	return err
}

func readDataFromDB(bucketName string, db *bolt.DB, keys []string) error {
	utils.Title("Reading data from bucket")
	err := db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("bucket:%q does not exists", bucketName)
		}

		for _, k := range keys {
			v := b.Get([]byte(k))
			if v == nil {
				logger.Warnf("No value found for key:%q in bucket:%q", k, bucketName)
				continue
			}
			logger.Debugf("Value found for key:%q=%v", k, utils.BytesToInt(v))
		}

		return nil
	})
	return err
}

func showAllData(bucketName string, db *bolt.DB, keyParser ByteToStringConverter, valueParser ByteToStringConverter) error {
	utils.Title("Showing all data from bucket")
	if keyParser == nil {
		keyParser = utils.BytesToString
	}
	if valueParser == nil {
		valueParser = utils.BytesToString
	}

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		if b == nil {
			return fmt.Errorf("bucket:%q does not exists", bucketName)
		}

		cursor := b.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			logger.Debugf("Key=%v, value=%v", keyParser(k), valueParser(v))
		}

		return nil
	})
	return err
}

type User struct {
	ID        int    `json:"int"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func runUserExample(db *bolt.DB) {
	utils.Title("Running User example")
	userBucketName := "USERS"

	showAllData(
		userBucketName,
		db,
		func(b []byte) string {
			return strconv.Itoa(utils.BytesToInt(b))
		},
		func(b []byte) string {
			user := User{}
			err := json.Unmarshal(b, &user)
			if err != nil {
				return fmt.Sprintf("Error while unmarshalling: %v", err)
			}
			return fmt.Sprintf("%+v", user)
		},
	)
	// createBucket(userBucketName, db)

	// users := []*User{
	// 	{FirstName: "John", LastName: "Doe"},
	// 	{FirstName: "Jane", LastName: "Doe"},
	// 	{FirstName: "John", LastName: "Smith"},
	// 	{FirstName: "Jane", LastName: "Smith"},
	// }

	// for _, u := range users {
	// 	err := createUser(userBucketName, db, u)
	// 	if err != nil {
	// 		logger.Errorf("error while adding user:%v to db.%s, %v", u, userBucketName, err)
	// 		continue
	// 	}
	// 	logger.Debugf("Successfully added user:%v to db.%s with id:%d", u, userBucketName, u.ID)
	// }

	// for _, u := range users {
	// 	fmt.Printf("u.ID: %v\n", u.ID)
	// }
}

func createUser(userBucket string, db *bolt.DB, user *User) error {
	return db.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(userBucket))
		if b == nil {
			return fmt.Errorf("%q bucket not found", userBucket)
		}

		id, _ := b.NextSequence()
		user.ID = int(id)

		buf, err := json.Marshal(user)
		if err != nil {
			return err
		}

		return b.Put(utils.IntToBytes(user.ID), buf)
	})
}
