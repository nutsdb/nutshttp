package nutshttp

import "github.com/xujiajun/nutsdb"

// Update handle insert and update operation
func (c *core) Update(bucket string, key string, value string, ttl uint32) error {
	err := c.db.Update(func(tx *nutsdb.Tx) error {
		err := tx.Put(bucket, []byte(key), []byte(value), ttl)
		return err
	})
	return err
}

// Delete handle delete operation
func (c *core) Delete(bucket string, key string) error {
	err := c.db.Update(func(tx *nutsdb.Tx) error {
		err := tx.Delete(bucket, []byte(key))
		return err
	})
	return err
}

// Get handle get key operation
func (c *core) Get(bucket string, key string) (value string, err error) {
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

func (c *core) PrefixScan(bucket string, prefix string, offSet int, limNum int) (entries nutsdb.Entries, err error) {
	err = c.db.View(func(tx *nutsdb.Tx) error {
		entries, _, err = tx.PrefixScan(bucket, []byte(prefix), offSet, limNum)
		return err
	})
	return entries, err
}

func (c *core) PrefixSearchScan(bucket, prefix string, reg string, offSet int, limNum int) (entries nutsdb.Entries, err error) {
	err = c.db.View(func(tx *nutsdb.Tx) error {
		entries, _, err = tx.PrefixSearchScan(bucket, []byte(prefix), reg, offSet, limNum)
		return err
	})
	return entries, err
}

func (c *core) RangeScan(bucket string, start string, end string) (entries nutsdb.Entries, err error) {
	err = c.db.View(func(tx *nutsdb.Tx) error {
		entries, err = tx.RangeScan(bucket, []byte(start), []byte(end))
		return err
	})
	return entries, err
}

func (c *core) GetAll(bucket string) (entries nutsdb.Entries, err error) {
	err = c.db.View(func(tx *nutsdb.Tx) error {
		entries, err = tx.GetAll(bucket)
		return err
	})
	return entries, err
}
