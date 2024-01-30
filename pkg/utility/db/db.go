package db

import (
	"github.com/nutsdb/nutsdb"
	"log"
	"palworld-chan/internal/consts"
	"sync"
)

var instance *nutsdb.DB
var once sync.Once

func createBucket(db *nutsdb.DB) {
	err := db.Update(
		func(tx *nutsdb.Tx) error {
			return tx.NewBucket(nutsdb.DataStructureBTree, consts.BUCKET)
		})
	if err != nutsdb.ErrBucketAlreadyExist {
		log.Fatal(err)
	}
}

func Db() *nutsdb.DB {
	once.Do(func() {
		localDb, err := nutsdb.Open(
			nutsdb.DefaultOptions,
			nutsdb.WithDir("/tmp/nutsdb"), // 数据库会自动创建这个目录文件
		)
		if err != nil {
			log.Fatal(err)
		}
		instance = localDb
		createBucket(instance)

	})
	return instance
}
