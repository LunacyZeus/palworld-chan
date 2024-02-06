package rcon

import (
	"fmt"
	"github.com/gorcon/rcon"
	"log"
	"palworld-chan/internal/consts"
	"palworld-chan/internal/service/api/models"
	"palworld-chan/pkg/logger"
	"palworld-chan/pkg/utility/utils"
	"strings"
	"time"
)

// 命令来自 https://tech.palworldgame.com/server-commands
// 部分代码借鉴于 https://github.com/Verthandii/palworld-go/blob/main/rcon/client.go

type CmdName string

const (
	Shutdown         CmdName = "Shutdown"         // /Shutdown {Seconds} {MessageText}
	DoExit           CmdName = "DoExit"           // /DoExit
	Broadcast        CmdName = "Broadcast"        // /Broadcast {MessageText}
	KickPlayer       CmdName = "KickPlayer"       // /KickPlayer {SteamID}
	BanPlayer        CmdName = "BanPlayer"        // /BanPlayer {SteamID}
	TeleportToPlayer CmdName = "TeleportToPlayer" // /TeleportToPlayer {SteamID}
	TeleportToMe     CmdName = "TeleportToMe"     // /TeleportToMe {SteamID}
	ShowPlayers      CmdName = "ShowPlayers ."    // /ShowPlayers
	Info             CmdName = "Info"             // /Info
	Save             CmdName = "Save"             // /Save
)

// Client .
type Client struct {
	Address       string
	AdminPassword string
	conn          *rcon.Conn
	skipErrors    bool
}

// New .
func New(Address, AdminPassword string) (*Client, error) {
	// 设置自定义超时时间
	dialTimeout := consts.RCON_TIMEOUT * time.Second
	deadline := consts.RCON_TIMEOUT * time.Second

	conn, err := rcon.Dial(Address, AdminPassword, rcon.SetDialTimeout(dialTimeout), rcon.SetDeadline(deadline))
	if err != nil {
		return nil, err
	}

	return &Client{
		Address:       Address,
		AdminPassword: AdminPassword,
		conn:          conn,
		skipErrors:    true,
	}, nil
}

// HandleMemoryUsage 发广播 重启维护
func (c *Client) HandleMemoryUsage(threshold float64) {
	c.Broadcast(fmt.Sprintf("now_server_load_Memory_is_above_%v%%", threshold))
	c.Broadcast("need_maintance")
	c.Save()
	c.Shutdown("60", "this_server_will_reboot_soon_due_to_memory_leak_after_60s")
}

// Shutdown The server is shut down after the number of Seconds Will be notified of your MessageText.
func (c *Client) Shutdown(seconds, message string) (string, error) {
	return c.exec(Shutdown, seconds, message)
}

// DoExit Force stop the server.
func (c *Client) DoExit() (string, error) {
	return c.exec(DoExit)
}

// Broadcast Send message to all player in the server.
func (c *Client) Broadcast(message string) (string, error) {
	return c.exec(Broadcast, message)
}

// KickPlayer Kick player from the server.
func (c *Client) KickPlayer(steamId string) (string, error) {
	return c.exec(KickPlayer, steamId)
}

// BanPlayer BAN player from the server.
func (c *Client) BanPlayer(steamId string) (string, error) {
	return c.exec(BanPlayer, steamId)
}

// TeleportToPlayer Teleport to current location of target player.
func (c *Client) TeleportToPlayer(steamId string) (string, error) {
	return c.exec(TeleportToPlayer, steamId)
}

// TeleportToMe Target player teleport to your current location
func (c *Client) TeleportToMe(steamId string) (string, error) {
	return c.exec(TeleportToMe, steamId)
}

// ShowPlayers Show information on all connected players.
func (c *Client) ShowPlayers() ([]models.OnlinePlayer, error) {
	response, err := c.execute(ShowPlayers)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	OnlinePlayers := []models.OnlinePlayer{}

	lines := strings.Split(response, "\n")
	titles := strings.Split(lines[0], ",")
	var result []map[string]string
	for _, line := range lines[1:] {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fields := strings.Split(line, ",")
		playerData := make(map[string]string)
		for i, title := range titles {
			value := "<null/err>"
			if i < len(fields) {
				value = fields[i]
				if strings.Contains(value, "\u0000") {
					// \u0000 直接显示null
					value = "<null/err>"
				}
			}
			playerData[title] = value
		}
		name, _ := playerData["name"]
		playeruid, _ := playerData["playeruid"]
		steamid, _ := playerData["steamid"]
		online := utils.NowTimeStamp()

		onlinePlayer := models.OnlinePlayer{
			Name:      name,
			PlayerUid: playeruid,
			SteamId:   steamid,
			Online:    float64(online),
		}

		result = append(result, playerData)
		OnlinePlayers = append(OnlinePlayers, onlinePlayer)
	}
	return OnlinePlayers, nil
}

// Info Show server information.
func (c *Client) Info() (response string, err error) {
	response, err = c.execute(Info)
	if err != nil {
		return "", err
	}
	defer c.Close()

	return
}

// Save Save the world data.
func (c *Client) Save() (string, error) {
	return c.exec(Save)
}

func (c *Client) exec(cmd CmdName, args ...string) (result string, err error) {
	argStr := strings.Join(args, " ")
	cmdStr := string(cmd)
	if argStr != "" {
		cmdStr = fmt.Sprintf("%s %s", cmd, argStr)
	}

	result, err = c.conn.Execute(cmdStr)
	if err != nil {
		log.Printf("execute [%s] error [%v]\n", cmdStr, err)
		return
	}
	logger.Info("result->%s", result)
	return
}

func (c *Client) execute(cmd CmdName) (string, error) {
	cmdStr := string(cmd)
	response, err := c.conn.Execute(cmdStr)
	response = strings.TrimSpace(response)
	if err != nil && c.skipErrors && response != "" {
		return response, nil
	}
	return response, err
}

func (e *Client) Close() error {
	if e.conn != nil {
		return e.conn.Close()
	}
	return nil
}
