package nutshttp

import (
	"github.com/xujiajun/nutsdb"
)

func (c *core) listSet(bucket string, setName string) ([][]byte, error) {

	var (
		data [][]byte
		err  error
	)

	err = c.db.View(func(tx *nutsdb.Tx) error {
		data, err = tx.SMembers(bucket, []byte(setName))
		return err
	})

	return data, err
}
