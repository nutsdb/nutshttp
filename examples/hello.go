package main

import (
	"log"

	"nutshttp"

	"github.com/xujiajun/nutsdb"
)

func main() {
	// Open the database located in the /tmp/nutsdb directory.
	// It will be created if it doesn't exist.
	opt := nutsdb.DefaultOptions
	opt.Dir = "/tmp/nutsdb"
	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	go func() {
		if err := nutshttp.Enable(db); err != nil {
			panic(err)
		}
	}()

	testDB(db)

	select {}
}

func testDB(db *nutsdb.DB) {
	var (
		bucket = "bucket001"

		key   = []byte("foo")
		value = []byte("bar")
	)

	if err := db.Update(func(tx *nutsdb.Tx) error {
		return tx.SAdd(bucket, key, value)

	}); err != nil {
		log.Fatal(err)
	}

	if err := db.View(func(tx *nutsdb.Tx) error {
		items, err := tx.SMembers(bucket, key)
		if err != nil {
			return err
		}

		for _, item := range items {
			log.Printf("item: %s", item)
		}
		return nil
	}); err != nil {

		log.Fatal(err)
	}
}
