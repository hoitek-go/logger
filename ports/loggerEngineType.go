package ports

import "github.com/hoitek-go/logger/engines"

type LoggerEngineType interface {
	*engines.LoggerEngineFile | *engines.LoggerEngineStdout | any
}
