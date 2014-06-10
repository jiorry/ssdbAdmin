package page

import (
	"github.com/jiorry/gos"
)

type Queue struct {
	gos.Page
}

func (p *Queue) Init() {
	SetupPage(&p.Page, "queue")
	p.Title = "queue"
}
