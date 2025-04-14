package log

import (
	"go.uber.org/zap"
)

func GetZapLogger() *zap.Logger {
	return zap.NewExample()
}
