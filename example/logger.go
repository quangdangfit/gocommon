package main

import (
	"github.com/quangdangfit/gocommon/logger"
)

func main() {
	logger.Initialize("production")
	logger.Info("This info log")
}
