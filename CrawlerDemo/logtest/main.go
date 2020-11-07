package main

import (
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("test log",
		zap.String("string", "s"),
		zap.Int("int", 2),
	)

	_ = testf("hello %d %v", 1.0, "s")
}

func testf(format string, a ...interface{}) error {
	return nil
}
