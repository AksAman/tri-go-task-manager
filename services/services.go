package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"text/tabwriter"

	"github.com/AksAman/tri/dblib"
	"github.com/AksAman/tri/models"
	"github.com/AksAman/tri/utils"
	"github.com/boltdb/bolt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	bucketName = []byte("TODO")
	db         *bolt.DB
	logger     *zap.SugaredLogger
)

func init() {
	utils.InitializeLogger("tri.log")
	logger = utils.Logger

	_ = initDB()
	dblib.GetOrCreateBucket(db, bucketName)
	_ = db.Close()
}

func getDBFilePath() string {
	viperDBFile := viper.GetString("dbfile")
	if viperDBFile == "" {
		userHome, _ := homedir.Dir()
		return filepath.Join(userHome, ".tridos.db")
	}
	viperDBFile, _ = homedir.Expand(viperDBFile)
	return viperDBFile
}

func initDB() error {
	var err error
	db, err = dblib.InitDB(getDBFilePath())
	if err != nil {
		logger.Panicf("Error while initializing DB: %v", err)
		fmt.Println("Something went wrong while initializing DB, exiting!")
		os.Exit(1)
	}
	return nil
}

func ReadItems() ([]models.Item, error) {
	dbFilepath := getDBFilePath()
	if !utils.DoesFileExists(dbFilepath) {
		return []models.Item{}, errors.New(dbFilepath + " doesn't exist!")
	}

	err := initDB()
	if err != nil {
		return []models.Item{}, err
	}
	defer db.Close()

	var items []models.Item
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			item := models.Item{}
			err := json.Unmarshal(v, &item)
			if err != nil {
				return err
			}
			items = append(items, item)
		}

		return nil
	})

	if err != nil {
		return []models.Item{}, err
	}

	return items, nil
}

func SaveItems(items []models.Item) error {
	err := initDB()
	if err != nil {
		return err
	}
	defer db.Close()

	logger.Debugln("Adding items to DB", len(items))

	err = db.Update(func(tx *bolt.Tx) error {
		// bucket := tx.Bucket(bucketName)
		bucket := tx.Bucket(bucketName)

		for _, item := range items {
			id64, _ := bucket.NextSequence()
			item.ID = int(id64)
			item.Position = int(id64)

			buf, err := json.Marshal(item)
			if err != nil {
				return err
			}
			key := utils.IntToBytes(item.ID)
			err = bucket.Put(key, buf)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func MarkItemDoneByID(id int) error {
	err := initDB()
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)

		key := utils.IntToBytes(id)
		itemBuf := bucket.Get(key)
		if itemBuf == nil {
			return fmt.Errorf("no Item with id:%d found", id)
		}
		item := models.Item{}
		err := json.Unmarshal(itemBuf, &item)
		if err != nil {
			return err
		}

		item.Done = true

		itemBuf, _ = json.Marshal(item)
		err = bucket.Put(key, itemBuf)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func ShowTridos(items []models.Item, filterCondition func(item models.Item) bool) {
	utils.Title("Your TriDos")
	var filteredItems []models.Item
	for _, item := range items {
		if filterCondition(item) {
			filteredItems = append(filteredItems, item)
		}
	}

	if len(filteredItems) == 0 {
		fmt.Println("No TODOs found!")
	}

	sort.Sort(models.ItemsByPri(filteredItems))

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	defer w.Flush()

	for _, item := range filteredItems {
		_, err := fmt.Fprint(w, item.PrettyItem())
		if err != nil {
			return
		}
	}
}
