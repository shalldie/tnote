package db

import (
	"bytes"
	"encoding/gob"
	"os"
	"path/filepath"
	"reflect"
	"regexp"

	"github.com/shalldie/gog/gs"
	"github.com/syndtr/goleveldb/leveldb"
)

var homeDir, _ = os.UserHomeDir()

var CONFIG_FILE_PATH = filepath.Join(homeDir, ".ttm.db")

func LoadDB() *leveldb.DB {
	db, err := leveldb.OpenFile(CONFIG_FILE_PATH, nil)
	if err != nil {
		panic(err)
	}

	return db
}

func Get(key string, sender any) []byte {
	db := LoadDB()
	defer db.Close()

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

func GetMany[T any](keys []string, senders []T) {
	db := LoadDB()
	defer db.Close()

	gs.ForEach(keys, func(key string, index int) {
		sender := senders[index]

		data, err := db.Get([]byte(key), nil)
		if err != nil {
			panic(err)
		}

		if reflect.ValueOf(sender).Kind() == reflect.Pointer {
			decode := gob.NewDecoder(bytes.NewBuffer(data))
			decode.Decode(sender)
		}

	})
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

	db.Put([]byte(key), buffer.Bytes(), nil)
}

func FindKeysLike(patterns ...string) []string {
	db := LoadDB()
	defer db.Close()

	keys := []string{}
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := string(iter.Key())

		matched := gs.Every(patterns, func(pattern string, index int) bool {
			subMatched, _ := regexp.MatchString(pattern, key)
			return subMatched
		})

		if matched {
			keys = append(keys, key)
		}
	}
	iter.Release()

	return keys
}

func FindByPattern(patterns ...string) map[string][]byte {
	db := LoadDB()
	defer db.Close()

	m := map[string][]byte{}

	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := string(iter.Key())
		value := iter.Value()

		matched := gs.Every(patterns, func(pattern string, index int) bool {
			subMatched, _ := regexp.MatchString(pattern, key)
			return subMatched
		})

		if matched {
			m[key] = gs.Copy(value)
		}
	}
	iter.Release()

	return m
}
