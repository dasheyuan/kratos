package docs

import (
	"github.com/gobuffalo/packr/v2"
)

//go:generate packr2
func Create() (docs []byte, err error) {
	box := packr.New("api", "../api")
	for _, name := range box.List() {
		if name == "api.swagger.json" {
			docs, _ = box.Find(name)
		}
	}
	return
}
