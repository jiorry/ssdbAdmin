package page

import (
	"github.com/jiorry/gos"
)

type Hset struct {
	gos.Page
}

func (p *Hset) Init() {
	SetupPage(&p.Page, "hset")
	p.Title = "hset "
}
