package page

import (
	"github.com/jiorry/gos"
)

type Key struct {
	gos.Page
}

func (p *Key) Init() {
	SetupPage(&p.Page, "key")
	p.Title = "keys"
}
