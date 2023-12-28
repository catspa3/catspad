package prefixmanager

import (
	"github.com/catspa3/catspad/infrastructure/logger"
	"github.com/catspa3/catspad/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
