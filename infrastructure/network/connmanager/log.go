package connmanager

import (
	"github.com/gordanet/gord/infrastructure/logger"
	"github.com/gordanet/gord/util/panics"
)

var log = logger.RegisterSubSystem("CMGR")
var spawn = panics.GoroutineWrapperFunc(log)
