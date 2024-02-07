package models

type GameServerInfoStruct struct {
	ProcessName   string `json:"processName"`
	ServerName    string `json:"serverName"`
	ServerVersion string `json:"serverVersion"`
	MemoryUsage   string `json:"memoryUsage"`
	CpuUsage      string `json:"cpuUsage"`
	UpTime        string `json:"upTime"`
	OnlineCount   int    `json:"onlineCount"`
	LastBackUp    string `json:"lastBackUp"`
}

type ServerInfoStruct struct {
	HostName         string `json:"hostName"`
	Os               string `json:"os"`
	UpTime           string `json:"upTime"`
	CpuModel         string `json:"cpuModel"`
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
	Name      string  `json:"name"`
	PlayerUid string  `json:"playeruid"`
	SteamId   string  `json:"steamid"`
	Online    float64 `json:"online"`
}

type UpdateServerSettingInput struct {
	ProcessName     string `json:"processName" validate:"required"`
	ExecutablePath  string `json:"executablePath" validate:"required"`
	MemoryThreshold string `json:"memoryThreshold" validate:"gte=1,lte=99"`
	CheckPeriod     string `json:"checkPeriod" validate:"required,number"`
	RestartDelay    string `json:"restartDelay" validate:"required,number"`
	RconAddress     string `json:"rconAddress"`
	RconPort        string `json:"rconPort"`
	RconPasswd      string `json:"rconPasswd"`
	SourceDir       string `json:"sourceDir" validate:"required"`
	DestDir         string `json:"destDir" validate:"required"`
	BackupTime      string `json:"backupTime" validate:"required,number"`
	BackupCount     string `json:"backupCount" validate:"required,number"`
	AccessToken     string `json:"accessToken"`
	SecretKey       string `json:"secretKey"`
	Bucket          string `json:"bucket"`
}

type BackUpFile struct {
	FileName string `json:"name"`
	Created  string `json:"created"`
}

type SaveFile struct {
	FileName string `json:"name"`
	Created  string `json:"created"`
}
