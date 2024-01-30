package main

import (
	"embed"
	"palworld-chan/cmd"
	"palworld-chan/pkg/utility/embeds"
)

// Embed a directory
//
//go:embed frontend/dist/*
var embedDir embed.FS

//go:embed frontend/*
var f embed.FS

func init() {
	embeds.EmbedDir = embedDir
}

func main() {
	cmd.Execute()
}
