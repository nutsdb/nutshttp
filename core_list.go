package nutshttp

import "github.com/xujiajun/nutsdb"

func (c *core) LRange(bucket string, key string, start, end int) ([]string, error) {
	var items []string

	err := c.db.Update(func(tx *nutsdb.Tx) (err error) {

		itemsBytes, err := tx.LRange(bucket, []byte(key), start, end)
		if err == nil {
			for _, itemByte := range itemsBytes {
				items = append(items, string(itemByte))
			}
		}
		return err
	})

	return items, err
}

func (c *core) RPush(bucket, key, value string) error {
	err := c.db.Update(func(tx *nutsdb.Tx) error {
		err := tx.RPush(bucket, []byte(key), []byte(value))
		return err
	})
	return err
}

func (c *core) LPush(bucket, key, value string) error {
	err := c.db.Update(func(tx *nutsdb.Tx) error {
		err := tx.LPush(bucket, []byte(key), []byte(value))
		return err
	})
	return err
}

func (c *core) RPop(bucket, key string) (value string, err error) {
	err = c.db.Update(func(tx *nutsdb.Tx) error {
		resByte, err := tx.RPop(bucket, []byte(key))
		if err == nil {
			value = string(resByte)
		}
		return err
	})
	return value, err
}

func (c *core) LPop(bucket, key string) (value string, err error) {
	err = c.db.Update(func(tx *nutsdb.Tx) error {
		resByte, err := tx.RPop(bucket, []byte(key))
		if err == nil {
			value = string(resByte)
		}
		return err
	})
	return value, err
}

func (c *core) RPeek(bucket, key string) (value string, err error) {
	err = c.db.View(func(tx *nutsdb.Tx) error {
		resByte, err := tx.RPeek(bucket, []byte(key))
		if err == nil {
			value = string(resByte)
		}
		return nil
	})
	return value, err
}

func (c *core) LPeek(bucket, key string) (value string, err error) {
	err = c.db.View(func(tx *nutsdb.Tx) error {
		resByte, err := tx.LPeek(bucket, []byte(key))
		if err == nil {
			value = string(resByte)
		}
		return nil
	})
	return value, err
}

func (c *core) Rem(bucket, key, value string, count int) (num int, err error) {
	err = c.db.Update(func(tx *nutsdb.Tx) error {
		num, err = tx.LRem(bucket, []byte(key), count, []byte(value))
		return err
	})
	return num, err
}

func (c *core) Set(bucket, key, value string, index int) (err error) {
	err = c.db.Update(func(tx *nutsdb.Tx) error {
		err = tx.LSet(bucket, []byte(key), index, []byte(value))
		return err
	})
	return err
}

func (c *core) LTrim(bucket, key string, start, end int) (err error) {
	err = c.db.View(func(tx *nutsdb.Tx) error {
		err = tx.LTrim(bucket, []byte(key), start, end)
		return err
	})
	return err
}

func (c *core) LSize(bucket, key string) (size int, err error) {
	err = c.db.View(func(tx *nutsdb.Tx) error {
		size, err = tx.LSize(bucket, []byte(key))
		return err
	})
	return size, err
}
