package store

import "sync"

var blacklist = map[string]bool{}
var mu sync.Mutex

func Add(token string) {
	mu.Lock()
	defer mu.Unlock()
	blacklist[token] = true
}

func IsBlacklisted(token string) bool {
	mu.Lock()
	defer mu.Unlock()
	return blacklist[token]
}
