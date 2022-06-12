package model

import (
  "errors"
	"time"
)

type LoginRecord struct {
	accountID  string
	loginTime  *time.Time
	logoutTime *time.Time
}

type LoginHistory struct {
  records []LoginRecord
}

var Log LoginHistory

func (log *LoginHistory) Add(accountID string, t time.Time) {
  log.records = append(log.records, LoginRecord{accountID, &t, nil})
}

func (log *LoginHistory) IsAuthorized(t time.Time) (error, string) {
  if len(log.records) == 0 {
    return errors.New("No login record found"), ""
  } else {
    return nil, log.records[len(log.records)-1].accountID
  }
}
