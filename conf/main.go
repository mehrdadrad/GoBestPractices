package main

import (
	"fmt"
	"github.com/mehrdadrad/GoBestPractices/conf/config"
)

func main() {
	var cfg config.Config
	cfg.LoadJSON("config.json")
	fmt.Println(cfg.Get("maxProc"))
	fmt.Println(cfg.Get("path.tmp.linux"))
}
