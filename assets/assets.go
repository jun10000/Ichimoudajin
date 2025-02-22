package assets

import (
	"embed"
	_ "image/jpeg"
	_ "image/png"
)

//go:embed *
var Assets embed.FS
