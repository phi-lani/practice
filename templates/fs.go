package templates

import "embed"

//go:embed templates/*.gohtml
var FS embed.FS
