package models

import "errors"

var (
	ErrLimitMessagesReached = errors.New("ERR_LIMIT_REACHED")
	ErrServerError          = errors.New("ERR_INTERNAL_ERROR")
	ErrInvalidDomain        = errors.New("ERR_INVALID_DOMAIN")
	ErrInvalidAPIKey        = errors.New("ERR_INVALID_APIKEY")
	ErrInvalidMessage       = errors.New("ERR_INVALID_MESSAGE")
	ErrUnknownProvider      = errors.New("ERR_UNKNOWN_PROVIDER")
	ErrInvalidProviders     = errors.New("ERR_INVALID_PROVIDERS")
	ErrInvalidPublicAPIKey  = errors.New("ERR_INVALID_PUBAPIKEY")
	ErrLimitDailyMessages   = errors.New("ERR_DAILY_LIMIT_APIKEY")
	ErrAllLimitsReached     = errors.New("ERR_ALL_LIMITS_REACHED")
	ErrInvalidSender        = errors.New("ERR_INVALID_SENDER")
	ErrInvalidContent       = errors.New("ERR_INVALID_CONTENT")
	ErrInvalidTemplate      = errors.New("ERR_INVALID_TEMPLATE")
	ErrInvalidAttachment    = errors.New("ERR_INVALID_ATTACHMENT")
	ErrInvalidRecipients    = errors.New("ERR_INVALID_RECIPIENTS")
)
