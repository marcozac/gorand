package gorand

// Main character sets.
const (
	LOWERCASE = "abcdefghijklmnopqrstuvwxyz"
	UPPERCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	DIGIT     = "0123456789"
	SPECIAL   = "~=+%^*/()[]{}/!@#$?|"
)

// Two character set combinations.
const (
	// Combination of lowercase and uppercase sets.
	LETTERS = LOWERCASE + UPPERCASE

	// Combination of lowercase and digit sets.
	LOWER_DIGIT = LOWERCASE + DIGIT

	// Combination of lowercase and special sets.
	LOWER_SPECIAL = LOWERCASE + SPECIAL

	// Combination of uppercase and digit sets.
	UPPER_DIGIT = UPPERCASE + DIGIT

	// Combination of uppercase and special sets.
	UPPER_SPECIAL = UPPERCASE + SPECIAL

	// Combination of digit and special sets.
	DIGIT_SPECIAL = DIGIT + SPECIAL
)

// Three character set combinations.
const (
	// Combination of lowercase, uppercase and digit sets.
	ALPHANUM = LETTERS + DIGIT

	// Combination of lowercase, uppercase and special sets.
	LETTERS_SPECIAL = LETTERS + SPECIAL

	// Combination of lowercase, digit and special sets.
	LOWER_DIGIT_SPECIAL = LOWER_DIGIT + SPECIAL

	// Combination of uppercase, digit and special sets.
	UPPER_DIGIT_SPECIAL = UPPER_DIGIT + SPECIAL
)

// Combination of lowercase, uppercase, digit and special sets.
const ALL_CHARS = ALPHANUM + SPECIAL
