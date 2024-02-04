package tests

import (
	"palworld-chan/internal/consts"
	"palworld-chan/internal/service/api/models"
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

func aTest_db_zset(t *testing.T) {
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

func Test_user_zset(t *testing.T) {
	//dao.GetUsersFromRcon()

	player := models.OnlinePlayer{
		Name:      "hoho",
		PlayerUid: "1",
		SteamId:   "100",
		Online:    float64(time.Now().Unix()),
	}
	err := dao.AddUser(player)
	if err != nil {
		t.Errorf("name(%s) %v", player.Name, err)
	}

	player = models.OnlinePlayer{
		Name:      "hoho",
		PlayerUid: "1",
		SteamId:   "100",
		Online:    float64(time.Now().Unix()),
	}
	err = dao.AddUser(player)
	if err != nil {
		t.Errorf("name(%s) %v", player.Name, err)
	}

	player = models.OnlinePlayer{
		Name:      "hoho1",
		PlayerUid: "12",
		SteamId:   "100",
		Online:    float64(time.Now().Unix()),
	}
	err = dao.AddUser(player)
	if err != nil {
		t.Errorf("name(%s) %v", player.Name, err)
	}

	count, err := dao.CountUser()
	if err != nil {
		t.Errorf("count %v", err)
	}
	t.Logf("Users count: %d", count)

	userList, err := dao.ListUser()
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("Users: %v", userList)
}
