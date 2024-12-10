package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"

	"business-platform/app/common/web/respmiddleware"
	"business-platform/app/exchange/internal/config"
	"business-platform/app/exchange/internal/handler"
	"business-platform/app/exchange/internal/svc"
)

var configFile = flag.String("f", "app/exchange/etc/exchange.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	logx.MustSetup(c.Log)
	defer func() {
		if err := logx.Close(); err != nil {
			logx.Errorf("Error closing logx: %v", err)
			return
		}
		logx.Info("success closing logx")
	}()

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	server.Use(respmiddleware.NewResponseMiddleware().Handle)

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	logx.Infof("Application started at %s:%d...\n", c.Host, c.Port)

	// Create a context with additional fields
	logCtx := logx.ContextWithFields(context.Background(), logx.Field("key1", "value1"), logx.Field("key2", "value2"))

	// Use the context with fields to log
	logx.WithContext(logCtx).Info("这是一条带有额外字段的日志")

	// Example error handling
	//var err error // Assuming err is defined somewhere in your code
	//if err != nil {
	//	logx.WithContext(logCtx).Error(err)
	//}

	server.Start()
}
