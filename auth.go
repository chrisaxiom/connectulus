package auth

import (
	"math/rand"
	"time"
)

func ValidatePassword(u, p string) bool {
	return u == p
}

func ConnectRemote() {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
}
