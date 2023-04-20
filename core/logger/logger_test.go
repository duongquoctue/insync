package logger

import (
	"context"
	"testing"
)

const Name = "Tue Duong"

func Test_LogInfo(t *testing.T) {
	InitLogger()

	log := NewLoggerCtx(context.Background())

	log.Info("this", "is", "some", "args")
	log.Infow("this is msg", "key1", "value1", "key2", "value2", "key3", "value3")
	log.Infof("My name is %v", Name)
	log.Infoln("one", "good", "string")
}
