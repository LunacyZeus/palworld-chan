package db

import (
	"github.com/nutsdb/nutsdb"
	"palworld-chan/internal/consts"
	"palworld-chan/pkg/logger"
	"sync"
)

var instance *nutsdb.DB
var once sync.Once

func createBucket(db *nutsdb.DB, bucket string) {
	if err := db.Update(func(tx *nutsdb.Tx) error {
		// you should call Bucket with data structure and the name of bucket first
		return tx.NewBucket(nutsdb.DataStructureBTree, bucket)
	}); err != nil {
		if err == nutsdb.ErrBucketAlreadyExist {
			return
		}
		logger.Fatal("bucket创建异常->%v", err)
		return
	}
	logger.Info("创建本地bucket: %s", bucket)
}

func createSortSetBucket(db *nutsdb.DB, bucket string) {
	if err := db.Update(func(tx *nutsdb.Tx) error {
		// you should call Bucket with data structure and the name of bucket first
		return tx.NewSortSetBucket(bucket)
	}); err != nil {
		if err == nutsdb.ErrBucketAlreadyExist {
			return
		}
		logger.Fatal("bucket创建异常->%v", err)
		return
	}
	logger.Info("创建本地bucket: %s", bucket)
}

func Db() *nutsdb.DB {
	once.Do(func() {
		localDb, err := nutsdb.Open(
			nutsdb.DefaultOptions,
			nutsdb.WithDir("/tmp/pal_db"), // 数据库会自动创建这个目录文件
		)
		if err != nil {
			logger.Error("打开本地持久层错误->%v", err)
			return
		}
		instance = localDb
		createBucket(instance, consts.BUCKET)
		createBucket(instance, consts.USER_BUCKET)
		createSortSetBucket(instance, consts.SORT_BUCKET)
	})
	return instance
}
