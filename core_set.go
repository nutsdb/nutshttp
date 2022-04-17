package nutshttp

import (
	"github.com/xujiajun/nutsdb"
)

func (c *core) listSet(bucket string, key string) (data []string, err error) {

	var dataByte [][]byte

	err = c.db.View(func(tx *nutsdb.Tx) error {

		if dataByte, err = tx.SMembers(bucket, []byte(key)); err != nil {

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
			if err := tx.SAdd(bucket, []byte(key), itemsByte...); err != nil {

				return err
			}
		}

		return nil
	})

	return err
}

func (c *core) sAreMembers(bucket string, key string, items ...string) (ok bool, err error) {

	if err = c.db.View(
		func(tx *nutsdb.Tx) error {

			if len(items) > 0 {
				itemsByte := make([][]byte, len(items))
				for k, v := range items {
					itemsByte[k] = []byte(v)
				}

				if ok, err = tx.SAreMembers(bucket, []byte(key), itemsByte...); err != nil {
					return err
				}

			}

			return nil
		}); err != nil {

		return
	}

	return
}

func (c *core) sIsMember(bucket string, key string, item string) (ok bool, err error) {

	if err = c.db.View(

		func(tx *nutsdb.Tx) error {
			ok, err = tx.SIsMember(bucket, []byte(key), []byte(item))
			if err != nil {
				return err
			}

			return nil
		}); err != nil {

		return
	}

	return
}

func (c *core) sCard(bucket string, key string) (num int, err error) {

	if err := c.db.View(

		func(tx *nutsdb.Tx) error {

			num, err = tx.SCard(bucket, []byte(key))
			if err != nil {
				return err
			}

			return nil
		}); err != nil {

		return 0, err
	}

	return num, err
}
