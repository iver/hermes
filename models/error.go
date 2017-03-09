package models

import "errors"

var (
	ErrLimitDailyMessages  = errors.New(`Limite de mensajes diarios alcanzado`)
	ErrLimitPerHourMesages = errors.New(`Limite de mensajes por hora  alcanzado`)
	ErrServerError         = errors.New(`Error inesperado`)
	ErrAPIKeyNotDefined    = errors.New(`Api key no definida`)
)
