package session

import (
	"testing"
	"time"
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
		userSession := UserSession{
			UserName: td.UserName,
			Password: td.Password,
		}
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
	userSession := UserSession{
		sessionIsAuthenticated: true,
	}
	// connect
	userSession.Connect()

	// verify connection is established eventually
	isNil := true
	for i := 0; i < 51; i++ {
		time.Sleep(10 * time.Millisecond)
		if userSession.connection() != nil {
			isNil = false
			break
		}
	}

	if isNil {
		t.Error("expected connection to not be nil after a period of time")
	}
}

func TestConnectInValid(t *testing.T) {
	// create a session that is not authenticated
	userSession := UserSession{
		sessionIsAuthenticated: false,
	}
	// connect
	userSession.Connect()

	// verify that connection is not established
	isNotNil := false
	for i := 0; i < 51; i++ {
		time.Sleep(10 * time.Millisecond)
		if userSession.connection() != nil {
			isNotNil = true
			break
		}
	}

	if isNotNil {
		t.Error("expected connection to be nil after a period of time")
	}
}
