package gorand

import (
	"errors"
	"testing"
)

func charSetCheck(i int, set string) error {
	switch i {
	case 1:
		if set != LOWERCASE {
			return errors.New("test characters set from options failed. expected LOWERCASE, received: " + set)
		}
	case 2:
		if set != UPPERCASE {
			return errors.New("test characters set from options failed. expected UPPERCASE, received: " + set)
		}
	case 3:
		if set != LETTERS {
			return errors.New("test characters set from options failed. expected LETTERS, received: " + set)
		}
	case 4:
		if set != DIGIT {
			return errors.New("test characters set from options failed. expected DIGIT, received: " + set)
		}
	case 5:
		if set != LOWER_DIGIT {
			return errors.New("test characters set from options failed. expected LOWER_DIGIT, received: " + set)
		}
	case 6:
		if set != UPPER_DIGIT {
			return errors.New("test characters set from options failed. expected UPPER_DIGIT, received: " + set)
		}
	case 7:
		if set != ALPHANUM {
			return errors.New("test characters set from options failed. expected ALPHANUM, received: " + set)
		}
	case 8:
		if set != SPECIAL {
			return errors.New("test characters set from options failed. expected SPECIAL, received: " + set)
		}
	case 9:
		if set != LOWER_SPECIAL {
			return errors.New("test characters set from options failed. expected LOWER_SPECIAL, received: " + set)
		}
	case 10:
		if set != UPPER_SPECIAL {
			return errors.New("test characters set from options failed. expected UPPER_SPECIAL, received: " + set)
		}
	case 11:
		if set != LETTERS_SPECIAL {
			return errors.New("test characters set from options failed. expected LETTERS_SPECIAL, received: " + set)
		}
	case 12:
		if set != DIGIT_SPECIAL {
			return errors.New("test characters set from options failed. expected DIGIT_SPECIAL, received: " + set)
		}
	case 13:
		if set != LOWER_DIGIT_SPECIAL {
			return errors.New(
				"test characters set from options failed. expected LOWER_DIGIT_SPECIAL, received: " + set,
			)
		}
	case 14:
		if set != UPPER_DIGIT_SPECIAL {
			return errors.New(
				"test characters set from options failed. expected UPPER_DIGIT_SPECIAL, received: " + set,
			)
		}
	case 15:
		if set != ALL_CHARS {
			return errors.New("test characters set from options failed. expected ALL_CHARS, received: " + set)
		}
	default:
		return errors.New(
			"test characters set from options failed. received an unknown option: check the test function",
		)
	}
	return nil
}

func TestGetCharSet(t *testing.T) {
	t.Parallel()
	for i := 1; i <= 15; i++ {
		if set, err := GetCharSet(NewStringOpt(i)); err != nil {
			t.Fatal(err)
		} else if err := charSetCheck(i, set); err != nil {
			t.Fatal(err)
		}
	}
}

func TestGetCharSetErr(t *testing.T) { // check unknown option error
	t.Parallel()
	if _, err := GetCharSet(100); err == nil {
		t.Fatal("expected error for unknown option, received nil")
	}
}

func BenchmarkCharSetFromOptAlphanum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := GetCharSet(WITH_LOWERCASE | WITH_UPPERCASE | WITH_DIGIT); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkCharSetFromOptAllSets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 1; i <= 15; i++ {
			if _, err := GetCharSet(NewStringOpt(i)); err != nil {
				b.Fatal(err)
			}
		}
	}
}
