package nutshttp

import "github.com/xujiajun/nutsdb"

// SUpdate handle insert and update operation
func (c *core) SUpdate(bucket string, key string, value string, ttl uint32) error {
	err := c.db.Update(func(tx *nutsdb.Tx) error {
		err := tx.Put(bucket, []byte(key), []byte(value), ttl)
		return err
	})
	return err
}

// SDelete handle delete operation
func (c *core) SDelete(bucket string, key string) error {
	err := c.db.Update(func(tx *nutsdb.Tx) error {
		err := tx.Delete(bucket, []byte(key))
		return err
	})
	return err
}

// SGet handle get key operation
func (c *core) SGet(bucket string, key string) (value string, err error) {
	err = c.db.View(func(tx *nutsdb.Tx) error {
		entry, err := tx.Get(bucket, []byte(key))
		if err != nil {
			return err
		}
		value = string(entry.Value)
		return nil
	})
	return value, err
}
