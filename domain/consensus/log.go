package consensus

import (
	"github.com/catspa3/catspad/infrastructure/logger"
	"github.com/catspa3/catspad/util/panics"
)

var log = logger.RegisterSubSystem("BDAG")
var spawn = panics.GoroutineWrapperFunc(log)
