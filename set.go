package nutshttp

import (
	"log"
	"net/http"

	"github.com/xujiajun/nutsdb"
)

func listSet(db *nutsdb.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("hello world")

		// vars := mux.Vars(r)
		// log.Printf("got: %+v", vars["bucket"])

		// if err := db.View(func(tx *nutsdb.Tx) error {
		// 	items, err := tx.SMembers(bucket, key)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	for _, item := range items {
		// 		log.Printf("item: %s", item)
		// 	}
		// 	return nil
		// }); err != nil {

		// 	log.Fatal(err)
		// }

	}
}

func makeSMembers(db *nutsdb.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("hello world")

	}
}
