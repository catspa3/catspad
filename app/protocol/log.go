package protocol

import (
	"github.com/catspa3/catspad/infrastructure/logger"
	"github.com/catspa3/catspad/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
