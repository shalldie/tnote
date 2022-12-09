// github.com/dgraph-io/badger/v3 更快，但是更大，，， 打包后多出来 6.5M

package db

import (
	"bytes"
	"encoding/gob"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"sync"

	"github.com/shalldie/gog/gs"
	"github.com/syndtr/goleveldb/leveldb"
)

var HomeDir, _ = os.UserHomeDir()

var CONFIG_FILE_PATH = filepath.Join(HomeDir, ".ttm.leveldb.db")

var dbm *sync.Mutex = &sync.Mutex{}

func LoadDB() *leveldb.DB {

	// opt := badger.DefaultOptions(CONFIG_FILE_PATH)
	// opt.Logger = nil
	// db, err := badger.Open(opt)

	// if err != nil {
	// 	panic(err)
	// }

	db, err := leveldb.OpenFile(CONFIG_FILE_PATH, nil)

	if err != nil {
		panic(err)
	}

	return db
}

func Get(key string, sender any) []byte {
	// dbm.Lock()
	// defer dbm.Unlock()
	db := LoadDB()
	defer db.Close()

	// var data []byte

	// db.View(func(txn *badger.Txn) error {
	// 	item, err := txn.Get([]byte(key))
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	item.Value(func(val []byte) error {
	// 		data = append([]byte{}, val...)
	// 		return nil
	// 	})
	// 	return nil
	// })

	data, err := db.Get([]byte(key), nil)

	if err != nil {
		panic(err)
	}

	if reflect.ValueOf(sender).Kind() == reflect.Pointer {
		decode := gob.NewDecoder(bytes.NewBuffer(data))
		decode.Decode(sender)
	}

	return data
}

func Save(key string, sender any) {
	dbm.Lock()
	defer dbm.Unlock()

	var buffer bytes.Buffer
	encode := gob.NewEncoder(&buffer)
	err := encode.Encode(sender)

	if err != nil {
		panic(err)
	}

	db := LoadDB()
	defer db.Close()

	// err = db.Update(func(txn *badger.Txn) error {
	// 	err := txn.Set([]byte(key), buffer.Bytes())
	// 	return err
	// })

	err = db.Put([]byte(key), buffer.Bytes(), nil)

	if err != nil {
		panic(err)
	}
}

func FindByPattern(patterns ...string) map[string][]byte {
	dbm.Lock()
	defer dbm.Unlock()

	db := LoadDB()
	defer db.Close()

	m := map[string][]byte{}

	// db.View(func(txn *badger.Txn) error {
	// 	it := txn.NewIterator(badger.DefaultIteratorOptions)
	// 	defer it.Close()

	// 	for it.Rewind(); it.Valid(); it.Next() {
	// 		item := it.Item()
	// 		key := string(item.Key())

	// 		matched := gs.Every(patterns, func(pattern string, index int) bool {
	// 			subMatched, _ := regexp.MatchString(pattern, key)
	// 			return subMatched
	// 		})

	// 		if !matched {
	// 			continue
	// 		}

	// 		item.Value(func(val []byte) error {
	// 			m[string(key)] = append([]byte{}, val...)
	// 			return nil
	// 		})

	// 	}
	// 	return nil
	// })

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		// Remember that the contents of the returned slice should not be modified, and
		// only valid until the next call to Next.
		key := string(iter.Key())
		value := iter.Value()

		matched := gs.Every(patterns, func(pattern string, index int) bool {
			subMatched, _ := regexp.MatchString(pattern, key)
			return subMatched
		})

		if matched {
			m[string(key)] = append([]byte{}, value...)
		}
	}
	iter.Release()

	return m
}

func Delete(key string) {
	dbm.Lock()
	defer dbm.Unlock()

	db := LoadDB()
	defer db.Close()

	db.Delete([]byte(key), nil)
}
