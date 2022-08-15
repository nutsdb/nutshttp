package main

import (
	"log"

	"nutshttp"

	"github.com/xujiajun/nutsdb"
)

func main() {
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

	// init test data
	initData(db)

	select {}
}
