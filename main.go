package main

import (
	_ "business-platform/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"business-platform/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
