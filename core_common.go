package nutshttp

import (
	"github.com/xujiajun/nutsdb"
)

// GetAllBuckets Regular obtains all buckets
func (c *core) GetAllBuckets(ds string, reg string) (buckets []string, err error) {
	if err = c.db.View(
		func(tx *nutsdb.Tx) error {
			var err error
			switch ds {
			case "string":
				err = tx.IterateBuckets(nutsdb.DataStructureBPTree, reg, func(bucket string) bool {
					buckets = append(buckets, bucket)
					return true
				})
			case "list":
				err = tx.IterateBuckets(nutsdb.DataStructureList, reg, func(bucket string) bool {
					buckets = append(buckets, bucket)
					return true
				})
			case "set":
				err = tx.IterateBuckets(nutsdb.DataStructureSet, reg, func(bucket string) bool {
					buckets = append(buckets, bucket)
					return true
				})
			case "zset":
				err = tx.IterateBuckets(nutsdb.DataStructureSortedSet, reg, func(bucket string) bool {
					buckets = append(buckets, bucket)
					return true
				})
			}
			return err
		},
	); err != nil {
		return nil, err
	}
	return buckets, nil
}
