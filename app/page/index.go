package page

import (
	"github.com/jiorry/gos"
)

type Index struct {
	gos.Page
}

func (p *Index) Init() {
	SetupPage(&p.Page, "index")
	p.Title = "ssdb admin"
	p.Layout.AddBottomRender(gos.NewPageTemplateRender("", "_footer", nil))
}

func SetupPage(p *gos.Page, name string) {
	p.JsPosition = "end"

	p.AddCss("site.css")
	p.AddCss("bootstrap.min.css")

	p.AddJs("jquery.js")
	p.AddJs("ajax.js")
	p.AddJs("site.js")

	data := make(map[string]bool)
	data["index"] = name == "index"
	data["key"] = name == "key"
	data["hset"] = name == "hset"
	data["zset"] = name == "zset"
	data["queue"] = name == "queue"

	p.Layout.AddTopRender(gos.NewPageTemplateRender("", "_header", data))
}
