package models

// PalSprint represents a character with various attributes.
type PalSprint struct {
	Hp        int64    `json:"hp"`        // Current health points of the character
	MaxHp     int64    `json:"maxHp"`     // Maximum health points of the character
	Gender    string   `json:"gender"`    // Gender of the character
	Type      string   `json:"type"`      // Type of the character
	Level     int32    `json:"level"`     // Level of the character
	IsBoss    bool     `json:"isBoss"`    // Indicates if the character is a boss
	Exp       int64    `json:"exp"`       // Experience points of the character
	Defense   int32    `json:"defense"`   // Defense skill level of the character
	Rank      int32    `json:"rank"`      // Rank of the character
	Melee     int32    `json:"melee"`     // Melee combat skill level of the character
	IsTower   bool     `json:"isTower"`   // Indicates if the character is a tower
	Skills    []string `json:"skills"`    // List of skills possessed by the character
	IsLucky   bool     `json:"isLucky"`   // Indicates if the character is lucky
	Ranged    int32    `json:"ranged"`    // Ranged combat skill level of the character
	Workspeed int32    `json:"workspeed"` // Work speed attribute of the character
}
