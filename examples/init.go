package main

import (
	"encoding/base64"
	"fmt"
	"github.com/xujiajun/nutsdb"
	"gopkg.in/yaml.v2"
	"math/rand"
	"os"
)

type initValueType struct {
	Score  float64 `yaml:"score"`
	Str    string  `yaml:"str"`
	Base64 string  `yaml:"base64"`
}

type initDataType struct {
	Kv   map[string]map[string]initValueType   `yaml:"kv"`
	List map[string]map[string][]initValueType `yaml:"list"`
	Set  map[string]map[string][]initValueType `yaml:"set"`
	ZSet map[string]map[string][]initValueType `yaml:"zset"`
}

func initData(db *nutsdb.DB) {
	bytes, err := os.ReadFile("examples/init.yaml")
	if err != nil {
		panic(err)
	}
	var yamlData initDataType
	err = yaml.Unmarshal(bytes, &yamlData)
	if err != nil {
		panic(err)
	}

	// init kv data
	initKvData(db, &yamlData)

	// init list data
	initListData(db, &yamlData)

	// init set data
	initSetData(db, &yamlData)

	// init sorted set data
	initSortedSetData(db, &yamlData)
}

func initKvData(db *nutsdb.DB, yamlData *initDataType) {
	_ = db.Update(func(tx *nutsdb.Tx) error {
		// init data from yaml
		for bucket, kv := range yamlData.Kv {
			for key, value := range kv {
				_ = tx.Put(bucket, []byte(key), getValue(value), 86_400_000)
			}
		}
		// multi value
		for i := 0; i < 200; i++ {
			_ = tx.Put("kv-bucket", []byte("key"), []byte(fmt.Sprintf("value-%d", i)), 86_400_000)
		}
		// multi key
		for i := 0; i < 200; i++ {
			_ = tx.Put("kv-bucket", []byte(fmt.Sprintf("key-%d", i)), []byte("value"), 86_400_000)
		}
		// multi bucket
		for i := 0; i < 200; i++ {
			_ = tx.Put(fmt.Sprintf("kv-bucket-%d", i), []byte("key"), []byte("value"), 86_400_000)
		}
		return nil
	})
}

func initListData(db *nutsdb.DB, yamlData *initDataType) {
	_ = db.Update(func(tx *nutsdb.Tx) error {
		// init data from yaml
		for bucket, kv := range yamlData.List {
			for key, values := range kv {
				for _, value := range values {
					_ = tx.RPush(bucket, []byte(key), getValue(value))
				}
			}
		}
		// multi value
		for i := 0; i < 200; i++ {
			_ = tx.RPush("list-bucket", []byte("key"), []byte(fmt.Sprintf("value-%d", i)))
		}
		// multi key
		for i := 0; i < 200; i++ {
			_ = tx.RPush("list-bucket", []byte(fmt.Sprintf("key-%d", i)), []byte("value"))
		}
		// multi bucket
		for i := 0; i < 200; i++ {
			_ = tx.RPush(fmt.Sprintf("list-bucket-%d", i), []byte("key"), []byte("value"))
		}
		return nil
	})
}

func initSetData(db *nutsdb.DB, yamlData *initDataType) {
	_ = db.Update(func(tx *nutsdb.Tx) error {
		// init data from yaml
		for bucket, kv := range yamlData.Set {
			for key, values := range kv {
				for _, value := range values {
					_ = tx.SAdd(bucket, []byte(key), getValue(value))
				}
			}
		}
		// multi value
		for i := 0; i < 200; i++ {
			_ = tx.SAdd("set-bucket", []byte("key"), []byte(fmt.Sprintf("value-%d", i)))
		}
		// multi key
		for i := 0; i < 200; i++ {
			_ = tx.SAdd("set-bucket", []byte(fmt.Sprintf("key-%d", i)), []byte("value"))
		}
		// multi bucket
		for i := 0; i < 200; i++ {
			_ = tx.SAdd(fmt.Sprintf("set-bucket-%d", i), []byte("key"), []byte("value"))
		}
		return nil
	})
}

func initSortedSetData(db *nutsdb.DB, yamlData *initDataType) {
	_ = db.Update(func(tx *nutsdb.Tx) error {
		// init data from yaml
		for bucket, kv := range yamlData.ZSet {
			for key, values := range kv {
				for _, value := range values {
					_ = tx.ZAdd(bucket, []byte(key), value.Score, getValue(value))
				}
			}
		}
		base := float64(rand.Intn(500))
		// multi value
		for i := 0; i < 200; i++ {
			_ = tx.ZAdd("sorted-set-bucket", []byte("key"), base+float64(i), []byte(fmt.Sprintf("value-%d", i)))
		}
		// multi key
		for i := 0; i < 200; i++ {
			_ = tx.ZAdd("sorted-set-bucket", []byte(fmt.Sprintf("key-%d", i)), base+float64(i), []byte("value"))
		}
		// multi bucket
		for i := 0; i < 200; i++ {
			_ = tx.ZAdd(fmt.Sprintf("sorted-set-bucket-%d", i), []byte("key"), base+float64(i), []byte("value"))
		}
		return nil
	})
}

func getValue(value initValueType) []byte {
	if len(value.Str) != 0 {
		return []byte(value.Str)
	} else if len(value.Base64) != 0 {
		bytes, err := base64.StdEncoding.DecodeString(value.Base64)
		if err != nil {
			panic(err)
		}
		return bytes
	}
	return []byte{}
}
