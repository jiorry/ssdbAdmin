package page

import (
	"github.com/jiorry/gos"
)

type ZsetItem struct {
	gos.Page
}

func (p *ZsetItem) Init() {
	SetupPage(&p.Page, "")
	name := p.Ctx.FormValue("q")
	p.SetData(name)
}
