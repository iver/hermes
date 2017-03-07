package models

package models

import "errors"

var (
	LimitDailyMessages  = errors.New(`Limite de mensajes diarios alcanzado`)
	LimitPerHourMesages = errors.New(`Limite de mensajes por hora  alcanzado`)
	ServerError         = errors.New(`Error inesperado`)
	ApiKeyNotDefined    = errors.New(`Api key no definida`)
)
