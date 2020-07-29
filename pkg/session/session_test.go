package session

import (
	"testing"
)

var authenticateTestCases = []struct {
	UserName string
	Password string
	Result   bool
}{
	{
		"hello",
		"world",
		false,
	},
	{
		"hello",
		"hello",
		true,
	},
	{
		"",
		"",
		true,
	},
	{
		"",
		"world",
		false,
	},
	{
		"hello",
		"",
		false,
	},
}

func TestAuthenticate(t *testing.T) {

	for _, td := range authenticateTestCases {
		result := Authenticate(td.UserName, td.Password)
		if result != td.Result {
			t.Errorf("expected %t, got %t", td.Result, result)
		}
	}
}

func TestUserSessionAuthenticate(t *testing.T) {
	for _, td := range authenticateTestCases {
		userSession := NewUserSession(
			WithUserName(td.UserName),
			WithPassword(td.Password))
		result := userSession.AuthenticateSession()
		if result != td.Result {
			t.Errorf("expected %t, got %t", td.Result, result)
		}
		authenticated := userSession.Authenticated()
		if authenticated != result {
			t.Errorf("expected %t, got %t", td.Result, authenticated)
		}
	}
}

func TestConnectValid(t *testing.T) {
	// create a session that is authenticated
	// may look to DRYify this instantiation, as it is now
	// very deep into the internal workings
	userSession := UserSession{
		sessionIsAuthenticated: true,
		connChan:               make(chan bool, 1),
	}

	// connect
	ch := userSession.Connect()
	connected := <-ch

	// expect connected to be true
	if !connected {
		t.Error("expected connected to be true")
	}

	// verify actual connection
	con := userSession.connection()
	if con == nil {
		t.Error("expected connection to be not nil")
	}
}

func TestConnectInValid(t *testing.T) {
	// create a session that
	userSession := UserSession{
		sessionIsAuthenticated: false,
		connChan:               make(chan bool, 1),
	}

	// connect
	ch := userSession.Connect()
	connected := <-ch

	// expect connected to be false
	if connected {
		t.Error("expected connected to be false")
	}

	// verify actual connection is nil
	con := userSession.connection()
	if con != nil {
		t.Error("expected connection to be nil")
	}
}
