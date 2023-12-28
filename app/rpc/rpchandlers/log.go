package rpchandlers

import (
	"github.com/catspa3/catspad/infrastructure/logger"
	"github.com/catspa3/catspad/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
