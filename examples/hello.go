package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"nutshttp"

	"github.com/xujiajun/nutsdb"
)

func main() {
	opt := nutsdb.DefaultOptions
	opt.Dir = "/tmp/nutsdb"

	files, _ := ioutil.ReadDir(opt.Dir)
	for _, f := range files {
		name := f.Name()
		if name != "" {
			err := os.RemoveAll(opt.Dir + "/" + name)
			if err != nil {
				panic(err)
			}
		}
	}

	db, err := nutsdb.Open(opt)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		db.Close()
		filepath.Clean(opt.Dir)
	}()

	go func() {
		if err := nutshttp.Enable(db); err != nil {
			panic(err)
		}
	}()

	// init test data
	initData(db)
	select {}
}
