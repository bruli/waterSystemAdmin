package templates

import (
	"embed"
)

//go:embed **/*.html *.html static/*
var FS embed.FS
