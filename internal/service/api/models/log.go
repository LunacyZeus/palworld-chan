package models

type LogResult struct {
	Lines          []string `json:"lines"`
	LastModifiedAt string   `json:"lastModifiedAt"`
}
