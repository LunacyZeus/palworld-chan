package dao

import (
	"fmt"
	"github.com/nutsdb/nutsdb"
	"palworld-chan/pkg/utility/db"
)

func Set(bucket, key, value string) (err error) {
	err = db.Db().Update(
		func(tx *nutsdb.Tx) error {
			bKey := []byte(key)
			bVal := []byte(value)
			if err = tx.Put(bucket, bKey, bVal, 0); err != nil {
				return err
			}
			return nil
		})
	if err != nil {
		return
	}
	return
}

func Get(bucket, key string) (value string, err error) {
	err = db.Db().View(
		func(tx *nutsdb.Tx) error {
			bKey := []byte(key)
			if bValue, err := tx.Get(bucket, bKey); err != nil {
				return err
			} else {
				value = string(bValue)
			}
			return nil
		})
	if err != nil {
		return
	}
	return
}

func Delete(bucket, key string) (err error) {
	err = db.Db().Update(
		func(tx *nutsdb.Tx) error {
			bKey := []byte(key)
			if err = tx.Delete(bucket, bKey); err != nil {
				return err
			}
			return nil
		})
	if err != nil {
		return
	}
	return

}

type member struct {
	score float64
	value string
}

func ZAdd(bucket, key, value string, score float64) (err error) {
	err = db.Db().Update(
		func(tx *nutsdb.Tx) error {
			bKey := []byte(key)
			return tx.ZAdd(bucket, bKey, score, []byte(value))
		})
	if err != nil {
		return
	}
	return
}

func ZCard(bucket, key string) (count int, err error) {
	err = db.Db().View(
		func(tx *nutsdb.Tx) error {
			bKey := []byte(key)
			if count, err = tx.ZCard(bucket, bKey); err != nil {
				return err
			}
			return nil
		})
	if err != nil {
		return
	}
	return
}

func ZMembers(bucket, key string) (members []member, err error) {
	err = db.Db().View(
		func(tx *nutsdb.Tx) error {
			bKey := []byte(key)
			if nodes, err := tx.ZMembers(bucket, bKey); err != nil {
				return err
			} else {
				fmt.Println("ZMembers:", nodes)

				for node, _ := range nodes {
					members = append(members, member{
						score: node.Score,
						value: string(node.Value),
					})
				}
			}
			return nil
		})

	if err != nil {
		return
	}
	return
}
