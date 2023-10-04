package main

import (
	"github.com/dressc-go/decryptorservice/pkg/config"
	"github.com/dressc-go/decryptorservice/pkg/server"
	"github.com/dressc-go/zlogger"
	"github.com/pkg/errors"
)

func main() {
	zlogger.SetGlobalLevel(zlogger.DebugLevel)
	logger := zlogger.GetLogger("main")
	logger.Debug().Msg("starting up")

	cnf := new(config.Config)
	err := cnf.New()
	if err != nil {
		e := errors.Wrap(err, "No config found")
		logger.Fatal().Err(e).Msg("")
	}

	server.Start(cnf)
}
