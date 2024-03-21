package main

import (
	"github.com/gordanet/gord/infrastructure/logger"
	"github.com/gordanet/gord/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("APLG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
