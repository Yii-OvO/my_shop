package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	_ "my_shop/internal/packed"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"my_shop/internal/cmd"
	_ "my_shop/internal/logic"
)

func main() {
	cmd.Main.Run(gctx.New())
}
