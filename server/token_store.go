package main

import "sync"

type TokenStore struct {
	tokens map[string]int
	mu     sync.RWMutex
}

func NewTokenStore() *TokenStore {
	return &TokenStore{
		tokens: make(map[string]int),
	}
}

func (ts *TokenStore) AddToken(token string, id int) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	ts.tokens[token] = id
}

func (ts *TokenStore) RemoveToken(token string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()
	delete(ts.tokens, token)
}

func (ts *TokenStore) ValidateToken(token string) bool {
	ts.mu.RLock()
	defer ts.mu.RUnlock()
	_, exists := ts.tokens[token]
	return exists
}

var InMemoryTokenStore = NewTokenStore()