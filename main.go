package main

import (
	"os"
	"github.com/faradayfan/Pepper"
	"strconv"
	"fmt"
)

func main() {

	log := pepper.New()
	args := os.Args[1:]

	if len(args)  < 1 {
		fmt.Print("Missing command line argument(s)\n" +
				"Please provide filename(s)")
		return
	}

	log.Info(strconv.Itoa(len(args)))
	log.Info(args[0])
}
