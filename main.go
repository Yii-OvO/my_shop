package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	_ "my_shop/internal/packed"

	_ "my_shop/internal/logic"

	"my_shop/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
