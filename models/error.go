package models

import "errors"

var (
	ErrLimitDailyMessages  = errors.New(`provider: Daily messages limit reached`)
	ErrLimitPerHourMesages = errors.New(`provider: Per hour messages limit reached`)
	ErrServerError         = errors.New(`provider: Internal error`)
	ErrAPIKeyNotDefined    = errors.New(`provider: APIKEY not defined`)
)
