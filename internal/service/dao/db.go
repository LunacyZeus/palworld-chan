package dao

import (
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
