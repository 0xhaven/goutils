package main

import (
	"flag"
	"fmt"

	"github.com/jacobhaven/goutils/log"
)

var arg string

func init() {
	flag.StringVar(&arg, "arg", "default", "arg for testing")
}

func main() {
	flag.Parse()

	fmt.Println(arg)
	log.Debug("debug")
	log.Info("info")
	log.Warning("warning")
	log.Error("error")
	log.Critical("crtical")
	log.Fatal("fatal")
}
