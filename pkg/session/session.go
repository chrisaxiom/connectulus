package session

import (
	"sync"

	"github.com/chrisaxiom/connectulus/pkg/conn"
	"github.com/chrisaxiom/connectulus/pkg/crypt"
)

// UserSession encapsulates the users session and information required for authentication
type UserSession struct {
	UserName               string
	Password               string
	sessionIsAuthenticated bool
	// the mutex protects the connection variable
	mutex        sync.RWMutex
	sophonicConn *conn.SophonicConnection
}

// AuthenticateSession will perform authentication on the given user session object
func (s *UserSession) AuthenticateSession() bool {
	s.sessionIsAuthenticated = crypt.ValidatePassword(s.UserName, s.Password)
	return s.sessionIsAuthenticated
}

// Authenticated returns true if the user session has been authenticated
func (s *UserSession) Authenticated() bool {
	return s.sessionIsAuthenticated
}

// Connect establishes a sophonic connection if the session is authenticated,
// and does nothing otherwise
func (s *UserSession) Connect() {
	if s.sessionIsAuthenticated {
		go func() {
			s.mutex.Lock()
			defer s.mutex.Unlock()
			s.sophonicConn = conn.ConnectSophon()
		}()
	}
}

func (s *UserSession) connection() *conn.SophonicConnection {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.sophonicConn
}

// Authenticate should perform authentication
func Authenticate(username, password string) bool {
	return crypt.ValidatePassword(username, password)
}
