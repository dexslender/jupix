package main

import (
	"github.com/dexslender/plane/plane"
	"github.com/dexslender/plane/utils"
	"github.com/disgoorg/log"
)

func main() {
	l := log.New(log.Ltime | log.Lshortfile)
	c := utils.SetupConfig(l, ".")
	if c.GetBool("debug-log-level") {
		l.SetLevel(log.LevelDebug)
	}

	p := plane.New(l, c)

	p.SetupBot()

	p.StartNLock()
}
