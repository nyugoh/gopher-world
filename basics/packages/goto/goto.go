package UrlShortener

import (
	"fmt"
	"strings"
	"sync"
)

type UrlStore struct {
	urls map[string]string
	mu sync.RWMutex
}

var Store *UrlStore

// Returns a new store
func GenerateNewUrlStore() *UrlStore { return &UrlStore{urls: make(map[string]string)} }

// Adds a long url with it's short url to store
func (s *UrlStore) Set(shortUrl, longUrl string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if exists := s.urls[shortUrl]; exists != "" {
		return false
	}
	s.urls[shortUrl] = longUrl
	return true
}

// Returns the long url when given a short url
func (s *UrlStore) Get(shortUrl string) string {
	// set read lock, no update to happen
	if strings.TrimSpace(shortUrl) == "" {
		return ""
	}
	s.mu.RLock() // Can be acquired by many threads for just read
	defer s.mu.RUnlock()
	return s.urls[shortUrl]
}

// Takes a url and generates the short url
func (s *UrlStore) Put(url string) string {
	for {
		key := s.genKey()
		if s.Set(key, url) {
			return key
		}
	}
}

// Return a new key
func (s *UrlStore) genKey() string { return fmt.Sprintf("%d", s.Count()) }

// Return the count of urls in store
func (s *UrlStore) Count() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.urls)
}