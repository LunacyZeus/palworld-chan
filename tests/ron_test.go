package tests

import (
	"log"
	"palworld-chan/pkg/utility/rcon"
	"testing"
)

func Test_Rcon_Info(t *testing.T) {
	endpoint := "127.0.0.1:25575"
	password := "test1234"

	rconClient, err := rcon.New(endpoint, password)
	if err != nil {
		t.Fatal(err)
	}
	rconClient.Info()

	broadcast := "msg,reboot,in,60,s"
	rconClient.Broadcast(broadcast)
	result, err := rconClient.ShowPlayers()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result_%v", result)
	//rconClient.HandleMemoryUsage(70)
}

func aTest_Rcon_Show(t *testing.T) {
	endpoint := "1"
	password := "2"
	conn := rcon.NewPalRcon(endpoint, password)
	players, err := conn.GetPlayers()
	if err != nil {
		t.Fatal(err)
	}
	log.Println(players)
}
