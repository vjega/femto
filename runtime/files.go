//go:generate go run assets_generate.go

package runtime

import "github.com/Bios-Marcel/femto"

var Files = femto.NewRuntimeFiles(files)
