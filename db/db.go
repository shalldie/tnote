package db

import (
	"bytes"
	"encoding/gob"
	"os"
	"path/filepath"
	"reflect"
	"regexp"

	"github.com/dgraph-io/badger/v3"
	"github.com/shalldie/gog/gs"
)

var homeDir, _ = os.UserHomeDir()

var CONFIG_FILE_PATH = filepath.Join(homeDir, ".ttm.badger.db")

func LoadDB() *badger.DB {

	opt := badger.DefaultOptions(CONFIG_FILE_PATH)
	opt.Logger = nil
	db, err := badger.Open(opt)

	if err != nil {
		panic(err)
	}

	return db
}

func Get(key string, sender any) []byte {
	db := LoadDB()
	defer db.Close()

	var data []byte

	db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			panic(err)
		}

		item.Value(func(val []byte) error {
			data = append([]byte{}, val...)
			return nil
		})
		return nil
	})

	if reflect.ValueOf(sender).Kind() == reflect.Pointer {
		decode := gob.NewDecoder(bytes.NewBuffer(data))
		decode.Decode(sender)
	}

	return data
}

func Save(key string, sender any) {

	var buffer bytes.Buffer
	encode := gob.NewEncoder(&buffer)
	err := encode.Encode(sender)

	if err != nil {
		panic(err)
	}

	db := LoadDB()
	defer db.Close()

	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), buffer.Bytes())
		return err
	})

	if err != nil {
		panic(err)
	}
}

func FindByPattern(patterns ...string) map[string][]byte {
	db := LoadDB()
	defer db.Close()

	m := map[string][]byte{}

	db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			key := string(item.Key())

			matched := gs.Every(patterns, func(pattern string, index int) bool {
				subMatched, _ := regexp.MatchString(pattern, key)
				return subMatched
			})

			if !matched {
				continue
			}

			item.Value(func(val []byte) error {
				m[string(key)] = append([]byte{}, val...)
				return nil
			})

		}
		return nil
	})

	return m
}
