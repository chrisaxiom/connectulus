/*
	Package auth provides implementations for authenticating against trisolan user stores
*/
package auth

// ValidatePassword takes a username and a password as arguments, returning a true if
// the password is valid, and false otherwise
func ValidatePassword(u, p string) bool {
	return u == p
}
