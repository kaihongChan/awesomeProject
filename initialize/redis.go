package initialize

import (
	"awesomeProject/global"
	"fmt"
)

func RedisInit() {
	redisCfg := global.Config.RedisConf

	fmt.Println(redisCfg.Addr)
}
