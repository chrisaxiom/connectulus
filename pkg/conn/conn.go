/*
	Package conn handles sophonic connections to trisolan resources
*/
package conn

import (
	"math/rand"
	"time"
)

// RemoteConnection is a struct that represents a connection to remote resource
type SophonicConnection struct {
}

// ConnectSophon establishes a sophonic connection to a remote resource and returns a
// SophonicConnection that can be used for communication
func ConnectSophon() *SophonicConnection {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	return &SophonicConnection{}
}
