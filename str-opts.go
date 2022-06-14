package gorand

import (
	"fmt"
)

// StrOpt is the option that defines which built-in characters set
// to use for pseudo-random string generation.
type NewStringOpt byte

const (
	// Enables lowercase characters set.
	WITH_LOWERCASE NewStringOpt = 1

	// Enables uppercase characters set.
	WITH_UPPERCASE = WITH_LOWERCASE << 1

	// Enables digit characters set.
	WITH_DIGIT = WITH_UPPERCASE << 1

	// Enables special characters set.
	WITH_SPECIAL = WITH_DIGIT << 1

	// Disables the use of Hex method in NewString.
	NO_HEX = WITH_SPECIAL << 1
)

// GetCharSet returns a built-in characters set from its option code
// returning an error if the option is not recognized.
// For example, GetCharSet(5) or GetCharSet(WITH_LOWERCASE|WITH_DIGIT)
// returns the LOWER_DIGIT characters set.
func GetCharSet(o NewStringOpt) (string, error) {
	switch o {
	case 1:
		return LOWERCASE, nil
	case 2:
		return UPPERCASE, nil
	case 3:
		return LETTERS, nil
	case 4:
		return DIGIT, nil
	case 5:
		return LOWER_DIGIT, nil
	case 6:
		return UPPER_DIGIT, nil
	case 7:
		return ALPHANUM, nil
	case 8:
		return SPECIAL, nil
	case 9:
		return LOWER_SPECIAL, nil
	case 10:
		return UPPER_SPECIAL, nil
	case 11:
		return LETTERS_SPECIAL, nil
	case 12:
		return DIGIT_SPECIAL, nil
	case 13:
		return LOWER_DIGIT_SPECIAL, nil
	case 14:
		return UPPER_DIGIT_SPECIAL, nil
	case 15:
		return ALL_CHARS, nil
	default:
		return "", fmt.Errorf("failed to get the characters set. unknown option: %d", o)
	}
}
