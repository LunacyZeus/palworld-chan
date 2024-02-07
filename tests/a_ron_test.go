package tests

import (
	"github.com/willroberts/minecraft-client"
	"log"
	"palworld-chan/pkg/utility/rcon"
	"testing"
)

func Test_1Rcon_Info(t *testing.T) {
	endpoint := "pal.zixing.fun:25575"
	password := "hoho123"

	rconClient, err := rcon.New(endpoint, password)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(rconClient.ShowPlayers())
}

func Test_2Rcon_Info(t *testing.T) {
	endpoint := "pal.zixing.fun:25575"
	password := "hoho123"

	// Create a new client and connect to the server.
	client, err := minecraft.NewClient(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Send some commands.
	if err := client.Authenticate(password); err != nil {
		log.Fatal(err)
	}
	resp, err := client.SendCommand("ShowPlayers")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.Body)
}
