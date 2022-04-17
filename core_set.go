package nutshttp

import (
	"github.com/xujiajun/nutsdb"
)

func (c *core) listSet(bucket string, key string) (data []string, err error) {

	var dataByte [][]byte

	err = c.db.View(func(tx *nutsdb.Tx) error {

		dataByte, err = tx.SMembers(bucket, []byte(key))
		if err != nil {
			return err
		}

		data = make([]string, len(dataByte))
		for k, v := range dataByte {
			data[k] = string(v)
		}

		return err
	})

	return data, err
}

func (c *core) addSet(bucket string, key string, items ...string) error {
	err := c.db.Update(func(tx *nutsdb.Tx) error {

		if len(items) > 0 {
			itemsByte := make([][]byte, len(items))
			for k, v := range items {
				itemsByte[k] = []byte(v)
			}
			err := tx.SAdd(bucket, []byte(key), itemsByte...)
			return err
		}

		return nil
	})

	return err
}
