package main

import (
	"github.com/iMax-pp/utils"
)

var logger *utils.Logger

func main() {
	logger, _ = utils.NewLogger("test_from_param.log", utils.LEVEL_TRACE)
	test()
	logger.Close()

	logger, _ = utils.NewLoggerFromConfig("logger.properties")
	test()
	logger.Close()
}

func test() {
	logger.TraceBegin("test")

	logger.Info("Logger level:", logger.Level)

	logger.Trace("test", 0)
	logger.Tracef("%d test %d", 1, 2)
	logger.Debug("test", 0)
	logger.Debugf("%d test %d", 1, 2)
	logger.Info("test", 0)
	logger.Infof("%d test %d", 1, 2)
	logger.Warn("test", 0)
	logger.Warnf("%d test %d", 1, 2)
	logger.Error("test", 0)
	logger.Errorf("%d test %d", 1, 2)

	logger.TraceEnd("test")
}
