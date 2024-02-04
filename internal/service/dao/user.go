package dao

import (
	"github.com/nutsdb/nutsdb"
	"palworld-chan/pkg/utility/db"
)

// 添加用户
func AddUser(bucket, key, value string) (err error) {
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
