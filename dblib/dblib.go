package dblib

import (
	"github.com/AksAman/tri/utils"
	"github.com/boltdb/bolt"
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func init() {
	utils.InitializeLogger("tri-db.log")
	logger = utils.Logger
}

// InitDB Make sure to close the db after using it
func InitDB(dbName string) (*bolt.DB, error) {
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}

type ByteToStringConverter func([]byte) string

//func AddMapDataToBucket[K comparable, V any](db *bolt.DB, bucketName string, data map[K]V, keyMarshaller func(K) []byte, valueMarshaller func(V) []byte) error {
//	utils.Title("Adding data to bucket")
//	err := db.Update(func(tx *bolt.Tx) error {
//		b := tx.Bucket([]byte(bucketName))
//		if b == nil {
//			return fmt.Errorf("bucket:%q does not exists", bucketName)
//		}
//		for k, v := range data {
//			err := b.Put(keyMarshaller(k), valueMarshaller(v))
//			if err != nil {
//				return err
//			}
//			logger.Debugf("Added %q:%d", k, v)
//		}
//		return nil
//	})
//	return err
//}

func GetOrCreateBucket(db *bolt.DB, bucketName []byte) *bolt.Bucket {
	var bucket *bolt.Bucket
	err := db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(bucketName)
		if err != nil {
			return err
		}
		bucket = b
		return nil
	})

	if err != nil {
		logger.Fatal(err)
	}

	return bucket
}

//func ReadDataFromDB(db *bolt.DB, keys []string, bucketName string) error {
//	utils.Title("Reading data from bucket")
//	err := db.View(func(tx *bolt.Tx) error {
//
//		b := tx.Bucket([]byte(bucketName))
//		if b == nil {
//			return fmt.Errorf("bucket:%q does not exists", bucketName)
//		}
//
//		for _, k := range keys {
//			v := b.Get([]byte(k))
//			if v == nil {
//				logger.Warnf("No value found for key:%q in bucket:%q", k, bucketName)
//				continue
//			}
//			logger.Debugf("Value found for key:%q=%v", k, utils.BytesToInt(v))
//		}
//
//		return nil
//	})
//	return err
//}

//func ShowAllData(db *bolt.DB, bucketName string, keyParser ByteToStringConverter, valueParser ByteToStringConverter) error {
//	utils.Title("Showing all data from bucket")
//	if keyParser == nil {
//		keyParser = utils.BytesToString
//	}
//	if valueParser == nil {
//		valueParser = utils.BytesToString
//	}
//
//	err := db.View(func(tx *bolt.Tx) error {
//		b := tx.Bucket([]byte(bucketName))
//		if b == nil {
//			return fmt.Errorf("bucket:%q does not exists", bucketName)
//		}
//
//		cursor := b.Cursor()
//
//		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
//			logger.Debugf("Key=%v, value=%v", keyParser(k), valueParser(v))
//		}
//
//		return nil
//	})
//	return err
//}
