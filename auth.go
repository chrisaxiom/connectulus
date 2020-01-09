package auth

import (
	"math/rand"
	"time"
)

type RemoteConnection struct {
}

func ValidatePassword(u, p string) bool {
	return u == p
}

func ConnectRemote() *RemoteConnection {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return &RemoteConnection{}
}
