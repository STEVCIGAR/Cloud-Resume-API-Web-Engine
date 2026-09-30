package assets

import (
	"embed"
	"io/fs"
)

//go:embed static/*
var staticFS embed.FS

// GetStaticFS returns the sub-filesystem for the static files.
func GetStaticFS() (fs.FS, error) {
	return fs.Sub(staticFS, "static")
}
