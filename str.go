package gorand

import (
	"encoding/base64"
	"strings"
)

// PseudoRandString is the generated pseudo-random string type.
type PseudoRandString string

// ToLower is a wrapper for go built-in strings homonym method.
// It reassign all string characters to their lowercase, then returns
// the PseudoRandString.
func (s PseudoRandString) ToLower() PseudoRandString {
	s = PseudoRandString(strings.ToLower(string(s)))
	return s
}

// ToUpper is a wrapper for go built-in strings homonym method.
// It reassign all string characters to their uppercase, then returns
// the PseudoRandString.
func (s PseudoRandString) ToUpper() PseudoRandString {
	s = PseudoRandString(strings.ToUpper(string(s)))
	return s
}

// ToBase64 is a wrapper for base64 string encoding that reassign the
// generated string to its base64 encoded version and return it.
func (s PseudoRandString) ToBase64() PseudoRandString {
	s = PseudoRandString(base64.StdEncoding.EncodeToString([]byte(s)))
	return s
}

// String returns the string typed PseudoRandString.
func (s PseudoRandString) String() string {
	return string(s)
}
