package models

type GameServerInfoStruct struct {
	ProcessName   string `json:"processName"`
	ServerName    string `json:"serverName"`
	ServerVersion string `json:"serverVersion"`
	MemoryUsage   string `json:"memoryUsage"`
	CpuUsage      string `json:"cpuUsage"`
	UpTime        string `json:"upTime"`
	LastBackUp    string `json:"lastBackUp"`
}

type ServerInfoStruct struct {
	HostName         string `json:"hostName"`
	Os               string `json:"os"`
	UpTime           string `json:"upTime"`
	CpuUsage         string `json:"cpuUsage"`
	MemoryUsage      string `json:"memoryUsage"`
	Load             string `json:"load"`
	DiskUsage        string `json:"diskUsage"`
	Bandwith         string `json:"bandwith"`
	IsVirtualization string `json:"isVirtualization"`
}

type BroadCastInput struct {
	BroadCast string `json:"broadCast"`
}

type OnlinePlayer struct {
	Name      string `json:"name"`
	PlayerUid string `json:"playeruid"`
	SteamId   string `json:"steamid"`
}
