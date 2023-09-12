package lib

import (
	"errors"
	"time"
)

type RateLimiterRecord map[string]UserRecord

type UserRecord struct {
	count           int
	accessTimeStamp time.Time
}

func (r *RateLimiterRecord) ValidateIPRequest(ip string) error {
	userRecord := (*r)[ip]
	if userRecord == (UserRecord{}) {
		userRecord = UserRecord{
			count:           1,
			accessTimeStamp: time.Now(),
		}
		(*r)[ip] = userRecord
		return nil
	} else {
		switch {
		case time.Since(userRecord.accessTimeStamp) > time.Minute:
			userRecord.accessTimeStamp = time.Now()
			userRecord.count = 1
			(*r)[ip] = userRecord
			return nil
		case userRecord.count >= 3:
			return errors.New("rate limit exceeded")
		default:
			userRecord.count += 1
			(*r)[ip] = userRecord
			return nil
		}
	}
}
