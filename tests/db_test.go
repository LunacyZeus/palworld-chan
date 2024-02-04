package tests

import (
	"palworld-chan/internal/consts"
	"palworld-chan/internal/service/dao"
	"testing"
	"time"
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

func Test_db_zset(t *testing.T) {
	bucket := consts.USER_BUCKET
	key := "test_user"
	value := "hoho2|1234|12345"
	timeUnix := time.Now().Unix()
	err := dao.ZAdd(bucket, key, value, float64(timeUnix))
	if err != nil {
		t.Logf("key(%s) %v", key, err)
	}
	err = dao.ZAdd(bucket, key, "hoho1|1234|12345", float64(timeUnix))
	if err != nil {
		t.Logf("key(%s) %v", key, err)
	}
	count, err := dao.ZCard(bucket, key)
	if err != nil {
		t.Logf("key(%s) %v", key, err)
	}
	t.Logf("bucket(%s) key(%s) count(%d)", bucket, key, count)

	members, err := dao.ZMembers(bucket, key)
	if err != nil {
		t.Logf("key(%s) %v", key, err)
	}
	t.Logf("bucket(%s) key(%s) members(%v)", bucket, key, members)
}
