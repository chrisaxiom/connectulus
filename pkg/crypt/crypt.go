/*
	Package crypt provides cryptographic primitives for Trisolian domains.

	Note that these primitives ONLY apply to Trisolan domains, and are
	considered unsafe for non-Trisolian applications.
*/
package crypt

// ValidatePassword takes a username and a password as arguments, returning a true if
// the password is valid, and false otherwise. Note that in this domain this is reflected
// via identity, i.e., returns true iff username is equal to password
func ValidatePassword(u, p string) bool {
	return u == p
}
