package logger

import "github.com/hoitek-go/logger/ports"

type LoggerEngine[T ports.LoggerEngineType] struct {
	Types     []string
	Instances *[]T
}
