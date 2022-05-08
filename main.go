package main

import (
	"flag"

	"github.com/atmom/nodego/common"
	"github.com/atmom/nodego/config"
	"github.com/atmom/nodego/server"
)

var (
	configPath = flag.String("config", "", "The nodego server configuration path")
)

func main() {
	// parse command flags
	flag.Parse()

	// load configuration from json file.
	c := config.Load(*configPath)

	// init log.
	common.InitLog(c)

	// start nodego server.
	server.Start(c)
}
