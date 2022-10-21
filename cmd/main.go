package main

import (
	"flag"
	"os"

	"github.com/sado0823/go-kitx/tpl/internal/conf"
	"github.com/sado0823/go-kitx/tpl/internal/di"

	_ "go.uber.org/automaxprocs"

	"github.com/sado0823/go-kitx/config"
	"github.com/sado0823/go-kitx/config/file"
	"github.com/sado0823/go-kitx/kit/log"
)

// go build -ldflags "-X 'main.Version=x.y.z' -X 'main.Name=demo'"
var (
	base         = new(di.Base)
	flagConfFile string
	Name         string
	Version      string
)

func init() {
	flag.StringVar(&flagConfFile, "conf", "../configs/config.yaml", "config path, eg: -conf config.yaml")
	base.ID, _ = os.Hostname()
	base.Name = Name
	base.Version = Version
}

func main() {
	flag.Parse()

	logger := log.WithFields(log.GetGlobal(),
		"service.id", base.ID,
		"service.name", base.Name,
		"service.version", base.Version,
		//"trace.id", tracing.TraceID(),
		//"span.id", tracing.SpanID(),
	)

	c := config.New(
		config.WithReader(
			file.New(flagConfFile),
		),
	)

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := di.WireApp(base, bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
