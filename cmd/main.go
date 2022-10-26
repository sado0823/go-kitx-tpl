package main

import (
	"flag"
	stdLog "log"
	"os"

	"github.com/sado0823/go-kitx/kit/tracing"
	"github.com/sado0823/go-kitx/tpl/internal/conf"
	"github.com/sado0823/go-kitx/tpl/internal/di"

	_ "go.uber.org/automaxprocs"

	"github.com/sado0823/go-kitx/config"
	"github.com/sado0823/go-kitx/config/file"
	"github.com/sado0823/go-kitx/kit/log"
)

// go build -ldflags "-X 'main.Version=x.y.z' -X 'main.Name=demo'"
var (
	base          = new(di.Base)
	flagConfFile  string
	Name          string
	Version       string
	TraceEndpoint string
)

func init() {
	flag.StringVar(&flagConfFile, "conf", "../configs/config.yaml", "config path, eg: -conf config.yaml")
	flag.StringVar(&TraceEndpoint, "trace", "http://localhost:14268/api/traces", "trace endpoint")
	base.ID, _ = os.Hostname()
	base.Name = Name
	base.Version = Version
}

func main() {
	flag.Parse()

	if err := tracing.Init(TraceEndpoint, base.Name+base.ID); err != nil {
		panic(err)
	}

	//logger := log.WithFields(log.GetGlobal(),
	logger := log.WithFields(log.NewStd(stdLog.Writer()),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", base.ID,
		"service.name", base.Name,
		"service.version", base.Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
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
