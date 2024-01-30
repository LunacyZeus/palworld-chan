package tests

import (
	"palworld-chan/internal/consts"
	"palworld-chan/internal/service/dao"
	"testing"
)

func Test_db_crud(t *testing.T) {
	bucket := consts.BUCKET
	key := "k"
	value := "v"
	err := dao.Set(bucket, key, value)
	if err != nil {
		t.Fatal(err)
	}

	value1, err := dao.Get(bucket, key)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("bucket(%s)-key(%s)-->%s", bucket, key, value1)

	//delete
	err = dao.Delete(bucket, key)
	if err != nil {
		t.Logf("Delete key(%s) %v", key, err)
	}

	//not found
	_, err = dao.Get(bucket, key)
	if err != nil {
		t.Logf("key(%s) %v", key, err)
	}

}
