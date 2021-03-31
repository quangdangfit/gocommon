package example

import (
	"github.com/quangdangfit/gocommon/logger"
)

func Logger() {
	logger.Initialize("production")
	logger.Info("This info log")
}
