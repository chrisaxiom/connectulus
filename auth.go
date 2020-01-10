/*
	Package auth provides (mock) implementations of functions useful for connecting users to a remote application.
*/
package auth

import (
	"math/rand"
	"time"
)

// RemoteConnection is a struct that represents a network connection to remote application
type RemoteConnection struct {
}

// ValidatePassword takes a username and a password as arguments, returning a true if
// the password is valid, and false otherwise
func ValidatePassword(u, p string) bool {
	return u == p
}

// ConnectRemote establishes a network connection to a remote application and returns a
// RemoteConnection that can be used for communication
func ConnectRemote() *RemoteConnection {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return &RemoteConnection{}
}
