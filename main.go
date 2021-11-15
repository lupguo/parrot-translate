package main

import (
	"log"

	"github.com/lupguo/parrot-translate/cmd"
)

func main() {
	// 命令行解析
	if err := cmd.Execute(); err != nil {
		log.Fatalf("command line execute got err:%v", err)
	}
}
