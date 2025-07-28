package main

import (
	"math/rand"
	"strings"
	"time"
)

func randomString(size int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var b strings.Builder
	b.Grow(size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		b.WriteByte(charset[rand.Intn(len(charset))])
	}
	return b.String()
}
