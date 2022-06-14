package gorand

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"math/big"
)

// NewString generates a new pseudo-random string on l characters
// using crypto methods and built-in character sets.
// For example, NewString(10, WITH_LOWERCASE|WITH_DIGIT) returns a
// pseudo-random string of length 10 using lowercase and digit
// character sets. Note that for digit + lowercase OR uppercase, NewHex
// method will be used instead of numbers generation to improve performances.
// This behavior can be disabled with NO_HEX option.
func NewString(l int, o NewStringOpt) (PseudoRandString, error) {
	if l <= 0 {
		return "", fmt.Errorf("pseudo-random string length can't be smaller than 1. Received: %d", l)
	}

	// use NewHex() generation for digit + (lowercase/uppercase) option
	switch o {
	case o & NO_HEX: // disable NewHex
	case 5:
		return NewHex(l)
	case 6:
		v, err := NewHex(l)
		return v.ToUpper(), err
	}

	charSet, err := GetCharSet(o)
	if err != nil {
		return "", err
	}

	b := make([]byte, l)
	for i := 0; i < l; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			return "", err
		}
		b[i] = charSet[n.Int64()]
	}
	return PseudoRandString(b), nil
}

// NewString generates a new pseudo-random string on l characters
// using crypto methods with a custom character set.
func NewStringWithSet(l int, charSet string) (PseudoRandString, error) {
	if l <= 0 {
		return "", fmt.Errorf("pseudo-random string length can't be smaller than 1. Received: %d", l)
	}

	b := make([]byte, l)
	for i := 0; i < l; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
		if err != nil {
			return "", err
		}
		b[i] = charSet[n.Int64()]
	}
	return PseudoRandString(b), nil
}

// NewHex returns a pseudo-random string combining crypto/rand bytes generation and hex encoding.
// This method is definitively faster than using by pseudo-random numbers generation.
func NewHex(l int) (PseudoRandString, error) {
	b := make([]byte, l)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return PseudoRandString(hex.EncodeToString(b)[:l]), nil
}
