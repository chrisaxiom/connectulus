package session

import (
	"sync"

	"github.com/chrisaxiom/connectulus/pkg/conn"
	"github.com/chrisaxiom/connectulus/pkg/crypt"
)

// UserSession encapsulates the users session and information required for authentication
// and connectivity with Sophonic
type UserSession struct {
	userName               string
	password               string
	sessionIsAuthenticated bool
	connChan               chan bool
	// the mutex protects the sophonicConn variable
	mutex        sync.RWMutex
	sophonicConn *conn.SophonicConnection
}

// UserSessionConfig encapsulates the required session configuration
type UserSessionConfig struct {
	UserName string
	Password string
}

// UserSessionConfigOption is a function to set configuration variables
type UserSessionConfigOption func(*UserSessionConfig)

// WithUserName sets the username
func WithUserName(username string) UserSessionConfigOption {
	return func(c *UserSessionConfig) {
		c.UserName = username
	}
}

// WithPassword sets the password
func WithPassword(password string) UserSessionConfigOption {
	return func(c *UserSessionConfig) {
		c.Password = password
	}
}

// NewUserSession creates a new user session given the config options
func NewUserSession(ops ...UserSessionConfigOption) *UserSession {
	conf := &UserSessionConfig{}

	for _, op := range ops {
		op(conf)
	}

	// TODO: some validation on config options

	return &UserSession{
		userName: conf.UserName,
		password: conf.Password,
		connChan: make(chan bool, 1),
	}
}

// Authenticatable allows for authentication of sessions
// without knowledge of the "finer details" of authentication methods.
// Might be useful for some future users of this library,
// however, I tend to only add interfaces when needed.
type Authenticatable interface {
	AuthenticateSession() bool
}

// Optional: declare a function type so this single-function interface
// can be implemented buy a function
// type AuthenticatableFunc func() bool

// func (f AuthenticatableFunc) AuthenticateSession() bool {
// 	return f()
// }

// UserSession is Authenticatable
var _ Authenticatable = (*UserSession)(nil)

// AuthenticateSession will perform authentication on the given user session object
func (s *UserSession) AuthenticateSession() bool {
	s.sessionIsAuthenticated = crypt.ValidatePassword(s.userName, s.password)
	return s.sessionIsAuthenticated
}

// AuthenticationCheckable allows for the checking of authentication status
// without knowledge of how authentication was performed
// Might be useful for some future users of this library,
// however, I tend to only add interfaces when needed.
type AuthenticationCheckable interface {
	Authenticated() bool
}

// Optional: declare a function type so this single-function interface
// can be implemented buy a function
// type AuthenticationCheckableFunc func() bool

// func (f AuthenticationCheckableFunc) Authenticated() bool {
// 	return f()
// }

// Authenticated returns true if the user session has been authenticated
func (s *UserSession) Authenticated() bool {
	return s.sessionIsAuthenticated
}

// UserSession is AuthenticationCheckable
var _ AuthenticationCheckable = (*UserSession)(nil)

// Connect establishes a sophonic connection if the session is authenticated
// Returns a channel that indicates if the connection was successful or not
// If the session is authenticated, then the channel returns true, and returns
// false otherwise
func (s *UserSession) Connect() chan bool {
	if s.sessionIsAuthenticated {
		go func() {
			s.mutex.Lock()
			s.sophonicConn = conn.ConnectSophon()
			s.mutex.Unlock()
			s.connChan <- true
		}()
	} else {
		s.connChan <- false
	}
	return s.connChan
}

// returns the connection for testing purposes
func (s *UserSession) connection() *conn.SophonicConnection {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.sophonicConn
}

// Authenticate performs authentication
func Authenticate(username, password string) bool {
	return crypt.ValidatePassword(username, password)
}
