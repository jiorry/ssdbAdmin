package page

import (
	"github.com/jiorry/gos"
)

type HsetItem struct {
	gos.Page
}

func (p *HsetItem) Init() {
	SetupPage(&p.Page, "")

	name := p.Ctx.FormValue("q")
	p.Title = "hset " + name
	p.SetData(name)
}
