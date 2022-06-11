package model

import (
	"time"
)

type login_history struct {
	AccountID  string
	LoginTime  *time.Time
	LogoutTime *time.Time
}
