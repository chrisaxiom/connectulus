package crypt

import "testing"

func TestValidatePasswordPass(t *testing.T) {
	if !ValidatePassword("alpha", "alpha") {
		t.Error("ValidatePassword should return true for equal values")
	}
}

func TestValidatePasswordFail(t *testing.T) {
	if ValidatePassword("alpha", "notalpha") {
		t.Error("ValidatePassword should return false for not-equal values")
	}
}
