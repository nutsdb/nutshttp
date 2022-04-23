package nutshttp

import (
	"github.com/xujiajun/nutsdb"
	"github.com/xujiajun/nutsdb/ds/zset"
)

func (c *core) zPut(bucket string, key []byte, score float64, value []byte) error {
	if err := c.db.Update(func(tx *nutsdb.Tx) error {
		if err := tx.ZAdd(bucket, key, score, value); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (c *core) zCard(bucket string) (num int, err error) {
	if err = c.db.View(func(tx *nutsdb.Tx) error {
		if num, err = tx.ZCard(bucket); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return num, err
	}
	return num, err
}

func (c *core) zCount(bucket string, start, end float64, options *zset.GetByScoreRangeOptions) (count int, err error) {
	if err = c.db.View(func(tx *nutsdb.Tx) error {
		if count, err = tx.ZCount(bucket, start, end, options); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return count, err
	}
	return count, nil
}

func (c *core) zGetByKey(bucket, key string) (node *zset.SortedSetNode, err error) {
	if err = c.db.View(func(tx *nutsdb.Tx) error {
		if node, err = tx.ZGetByKey(bucket, []byte(key)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return node, err
}

func (c *core) zMembers(bucket string) (nodes []*zset.SortedSetNode, err error) {
	var nodesMap map[string]*zset.SortedSetNode
	if err = c.db.View(func(tx *nutsdb.Tx) error {
		if nodesMap, err = tx.ZMembers(bucket); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	nodes = make([]*zset.SortedSetNode, len(nodesMap))
	for _, node := range nodesMap {
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func (c *core) zPeekMax(bucket string) (node *zset.SortedSetNode, err error) {
	if err = c.db.View(func(tx *nutsdb.Tx) error {
		if node, err = tx.ZPeekMax(bucket); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return node, nil
}

func (c *core) zPeekMin(bucket string) (node *zset.SortedSetNode, err error) {
	if err = c.db.View(func(tx *nutsdb.Tx) error {
		if node, err = tx.ZPeekMin(bucket); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return node, nil
}

func (c *core) zPopMax(bucket string) (node *zset.SortedSetNode, err error) {
	if err = c.db.View(func(tx *nutsdb.Tx) error {
		if node, err = tx.ZPopMax(bucket); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return node, nil
}

func (c *core) zPopMin(bucket string) (node *zset.SortedSetNode, err error) {
	if err = c.db.View(func(tx *nutsdb.Tx) error {
		if node, err = tx.ZPopMin(bucket); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return node, nil
}

func (c *core) zRangeByRank(bucket string, start, end int) (nodes []*zset.SortedSetNode, err error) {
	if err = c.db.View(func(tx *nutsdb.Tx) error {
		if nodes, err = tx.ZRangeByRank(bucket, start, end); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return nodes, nil
}

func (c *core) zRangeByScore(bucket string, start, end float64, options *zset.GetByScoreRangeOptions) (nodes []*zset.SortedSetNode, err error) {
	if err = c.db.View(func(tx *nutsdb.Tx) error {
		if nodes, err = tx.ZRangeByScore(bucket, start, end, options); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return nodes, nil
}

func (c *core) zRank(bucket, key string) (rank int, err error) {
	if err = c.db.View(func(tx *nutsdb.Tx) error {
		if rank, err = tx.ZRank(bucket, []byte(key)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return rank, err
	}
	return rank, nil

}

func (c *core) zRevRank(bucket, key string) (rank int, err error) {
	if err = c.db.View(func(tx *nutsdb.Tx) error {
		if rank, err = tx.ZRevRank(bucket, []byte(key)); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return rank, err
	}
	return rank, nil
}

func (c *core) zRem(bucket, key string) error {
	if err := c.db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.ZRem(bucket, key); err != nil {
				return err
			}
			return nil
		}); err != nil {
		return err
	}
	return nil
}

func (c *core) ZRemRangeByRank(bucket string, start, end int) error {
	if err := c.db.Update(
		func(tx *nutsdb.Tx) error {
			if err := tx.ZRemRangeByRank(bucket, start, end); err != nil {
				return err
			}
			return nil
		}); err != nil {
		return err
	}
	return nil
}

func (c *core) zScore(bucket, key string) (score float64, err error) {
	if err = c.db.View(
		func(tx *nutsdb.Tx) error {
			if score, err = tx.ZScore(bucket, []byte(key)); err != nil {
				return err
			}
			return nil
		}); err != nil {
		return 0, err
	}
	return score, nil
}
