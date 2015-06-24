package main

import (
	"flag"
	"github.com/mongodbinc-interns/mongoproxy"
	. "github.com/mongodbinc-interns/mongoproxy/log"
	"github.com/mongodbinc-interns/mongoproxy/modules"
	"github.com/mongodbinc-interns/mongoproxy/mongod"
)

var (
	port     int
	logLevel int
)

func parseFlags() {
	flag.IntVar(&port, "port", 8124, "port to listen on")
	flag.IntVar(&logLevel, "logLevel", DEBUG, "verbosity for logging")

	flag.Parse()
}

func main() {

	parseFlags()
	SetLogLevel(logLevel)

	// initialize the mockule
	module := mongod.MongodModule{}
	// initialize the pipeline
	chain := modules.CreateChain()
	modules.AddModule(chain, module)
	pipeline := modules.BuildPipeline(chain)

	mongoproxy.Start(port, pipeline)
}