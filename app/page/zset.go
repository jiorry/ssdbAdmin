package page

import (
	"github.com/jiorry/gos"
)

type Zset struct {
	gos.Page
}

func (p *Zset) Init() {
	SetupPage(&p.Page, "zset")
	p.Title = "zset"
}
