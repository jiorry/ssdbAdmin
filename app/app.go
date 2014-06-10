package main

import (
	"./api"
	"./page"
	"github.com/jiorry/gos"
	"github.com/jiorry/lib/ssdb"
)

func main() {
	gos.Init()

	gos.Route("/", (*page.Index)(nil))
	gos.Route("/key", (*page.Key)(nil))
	gos.Route("/hset", (*page.Hset)(nil))
	gos.Route("/hset/get", (*page.HsetItem)(nil))
	gos.Route("/zset", (*page.Zset)(nil))
	gos.Route("/zset/get", (*page.ZsetItem)(nil))
	gos.Route("/queue", (*page.Queue)(nil))

	// open api router
	gos.WebApiRoute("web", (*api.WebApi)(nil))

	ssdb.ConnectByConfig(gos.Configuration.GetConf("ssdb"))

	gos.Start()
}
