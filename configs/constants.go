package configs

import "time"

const (
	DefaultDateCsvLayoutFormat = "dd/mm/yyyy"
	AccessTokenExpireTime      = 15 * time.Minute
	RefreshTokenExpireTime     = 24 * time.Hour
	ReadHeaderTimeout          = 5 * time.Second
	TokenLifeSpan              = 24
)
