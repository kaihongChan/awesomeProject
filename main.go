package main

import (
	"awesomeProject/core"
	"awesomeProject/initialize"
)

func main() {
	initialize.ViperInit()
	core.RunServer()
}
