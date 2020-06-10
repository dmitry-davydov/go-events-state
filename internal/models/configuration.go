package models

import "context"

type Configuration struct {
	// порт http api сервера
	HttpServerPort int

	// время жизни события в статусе interrupted
	InterruptedValidPeriod int

	// сонтекст для завершения работы сервиса
	DieAppContext context.Context

	LogHttp bool
}
