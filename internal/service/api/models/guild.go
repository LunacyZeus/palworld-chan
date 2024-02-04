package models

//data from https://github.com/magicbear/palworld-server-toolkit

// GuildPlayer represents a player in a guild with basic information.
type GuildPlayer struct {
	PlayerUid string `json:"uid"`      // Unique identifier for the player
	Nickname  string `json:"nickname"` // Nickname of the player
}

// BasePlayer represents a player with concise information including stats and status.
type BasePlayer struct {
	PlayerUid      string           `json:"uid"`            // Unique identifier for the player
	Nickname       string           `json:"nickname"`       // Nickname of the player
	Level          int32            `json:"level"`          // Level of the player
	Hp             int64            `json:"hp"`             // Current health points of the player
	MaxHp          int64            `json:"maxHp"`          // Maximum health points of the player
	Exp            int64            `json:"exp"`            // Experience points of the player
	ShieldHp       int64            `json:"shieldHp"`       // Current shield health points of the player
	ShieldMaxHp    int64            `json:"shieldMaxHp"`    // Maximum shield health points of the player
	MaxStatusPoint int32            `json:"maxStatusPoint"` // Maximum status points the player can have
	StatusPoint    map[string]int32 `json:"statusPoint"`    // Map of different status points and their values
	FullStomach    float64          `json:"fullStomach"`    // Level of fullness of the player's stomach
}

// Player represents a player with detailed information, including BasePlayer details and a list of PalSprints.
type Player struct {
	BasePlayer              // Embedding BasePlayer struct for basic player information
	Pals       []*PalSprint `json:"pals"` // List of PalSprint characters associated with the player
}

// Guild represents a guild with its name, base camp level, admin player's UID, player list, and base IDs.
type Guild struct {
	Name           string         `json:"name"`            // Name of the guild
	AdminPlayerUid string         `json:"admin_uid"`       // UID of the guild's admin player
	BaseCampLevel  int32          `json:"base_camp_level"` // Level of the guild's base camp
	Players        []*GuildPlayer `json:"players"`         // List of players in the guild
	BaseIds        []string       `json:"base_ids"`        // List of base IDs associated with the guild
}
