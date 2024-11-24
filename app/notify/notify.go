package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"business-platform/app/notify/internal/config"
	"business-platform/app/notify/internal/handler"
	"business-platform/app/notify/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "app/notify/etc/notify.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.DisableStat()
	err := logx.SetUp(logx.LogConf{
		ServiceName: c.Name,
		Mode:        "console",
		Encoding:    "plain",
		TimeFormat:  "2006-01-02 15:04:05.000",
		Level:       "info",
		Path:        "logs",
	})
	if err != nil {
		return
	}

	server := rest.MustNewServer(c.RestConf,
		// enabled cors
		rest.WithCors(),
	)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
