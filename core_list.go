package nutshttp

import "github.com/xujiajun/nutsdb"

func (c *core) LRange(bucket string, key []byte, begin, offset int) ([][]byte, error) {
	var items [][]byte

	err := c.db.Update(func(tx *nutsdb.Tx) (err error) {

		items, err = tx.LRange(bucket, key, begin, offset)
		return err
	})

	return items, err
}
