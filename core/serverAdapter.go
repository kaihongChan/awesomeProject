package core

import (
	"awesomeProject/global"
	"awesomeProject/initialize"
	"fmt"
	"strconv"
	"time"
)

func RunServer() {
	Router := initialize.RouterInit()
	initialize.RedisInit()
	port := strconv.Itoa(global.Config.ServerConf.Port)
	s := initServer(":" + port, Router)
	time.Sleep(10 * time.Microsecond)
	fmt.Println("Server Start")
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}