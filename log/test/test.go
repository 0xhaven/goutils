package main

import (
	"flag"

	"github.com/jacobhaven/goutils/log"
)

var arg string

func init() {
	flag.StringVar(&arg, "arg", "default", "arg for testing")
}

func main() {
	flag.Parse()
	log.Debugf("debug: %s", arg)
	log.Infof("info: %s", arg)
	log.Warningf("warning: %s", arg)
	log.Errorf("error: %s", arg)
	log.Criticalf("critical: %s", arg)
	log.Fatalf("fatal: %s", arg)
}
