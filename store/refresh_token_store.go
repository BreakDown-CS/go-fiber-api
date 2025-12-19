package store

import (
	"time"
)

type RefreshToken struct {
	UserID int
	Expire time.Time
}

var refreshTokens = map[string]RefreshToken{}

func Save(token string, userID int, exp time.Time) {
	mu.Lock()

	defer mu.Unlock()

	refreshTokens[token] = RefreshToken{
		UserID: userID,
		Expire: exp,
	}
}

func Get(token string) (RefreshToken, bool) {
	mu.Lock()
	defer mu.Unlock()

	rt, ok := refreshTokens[token]
	return rt, ok
}

func Delete(token string) {
	mu.Lock()
	defer mu.Unlock()

	delete(refreshTokens, token)
}
