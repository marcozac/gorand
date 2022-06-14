package gorand

import (
	"encoding/base64"
	"testing"
)

const (
	strShortLen = 12
	strLongLen  = 120
)

func testNewStr(t *testing.T, strLen int) {
	t.Parallel()
	for i := 1; i <= 15; i++ {
		if s, err := NewString(strLen, NewStringOpt(i)); err != nil {
			t.Fail()
			t.Log("failed to generate pseudo-random string: ", err)
		} else if len(s) != strLen {
			t.Fail()
			t.Log("pseudo-random string generated with wrong length. wanted:", strLen, "received:", len(s))
		} else {
			t.Log("opt n:", i, "result:", s)
		}
	}
}

func TestNewStrShort(t *testing.T) {
	testNewStr(t, strShortLen)
}

func TestNewStrShortToLower(t *testing.T) {
	t.Parallel()
	if s, err := NewString(strShortLen, WITH_UPPERCASE|WITH_DIGIT|WITH_SPECIAL); err != nil {
		t.Fatal(err)
	} else {
		u := s.ToLower()
		t.Log("original str:", s, "after lower:", u)
	}
}
func TestNewStrShortToUpper(t *testing.T) {
	t.Parallel()
	if s, err := NewString(strShortLen, WITH_LOWERCASE|WITH_DIGIT|WITH_SPECIAL); err != nil {
		t.Fatal(err)
	} else {
		u := s.ToUpper()
		t.Log("original str:", s, "after upper:", u)
	}
}

func TestNewStrShortToBase64(t *testing.T) {
	t.Parallel()
	if s, err := NewString(strShortLen, WITH_LOWERCASE|WITH_DIGIT); err != nil {
		t.Fatal(err)
	} else if sb64, err := base64.StdEncoding.DecodeString(string(s.ToBase64())); err != nil {
		t.Fail()
		t.Log(err)
	} else if decSb64 := PseudoRandString(sb64[:]); decSb64 != s {
		t.Fail()
		t.Log(
			"base64 encoded string not matching the original one. wanted:", s,
			"received:", decSb64,
		)
	} else {
		t.Log("original str:", s, "after base64:", decSb64)
	}
}

func TestNewStringLengthErr(t *testing.T) {
	t.Parallel()
	if _, err := NewString(0, 1); err == nil {
		t.Fatal("expected string length < 1 error, received nil")
	}
}

func TestNewStringOptionErr(t *testing.T) {
	t.Parallel()
	if _, err := NewString(2, 100); err == nil {
		t.Fatal("expected unknown option error, received nil")
	}
}

func TestNewStrLong(t *testing.T) {
	testNewStr(t, strLongLen)
}

func TestNewStringCustomSet(t *testing.T) {
	t.Parallel()
	if s, err := NewStringWithSet(strShortLen, LOWERCASE); err != nil {
		t.Fatal(err)
	} else if len(s) != strShortLen {
		t.Fail()
		t.Log("pseudo-random string generated with wrong length. wanted:", strShortLen, "received:", len(s))
	} else {
		t.Log("with custom set; result:", s)
	}
}

func TestNewStringCustomSetLengthErr(t *testing.T) {
	t.Parallel()
	if _, err := NewStringWithSet(0, LOWERCASE); err == nil {
		t.Fatal("expected string length < 1 error, received nil")
	}
}

func BenchmarkNewStrShortLowercase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := NewString(strShortLen, WITH_LOWERCASE); err != nil {
			b.Fatal("failed to generate pseudo-random string: ", err)
		}
	}
}

func BenchmarkNewStrShortDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := NewString(strShortLen, WITH_DIGIT); err != nil {
			b.Fatal("failed to generate pseudo-random string: ", err)
		}
	}
}

func BenchmarkNewStrShortAlphanum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := NewString(strShortLen, WITH_LOWERCASE|WITH_UPPERCASE|WITH_DIGIT); err != nil {
			b.Fatal("failed to generate pseudo-random string: ", err)
		}
	}
}

func BenchmarkNewStrShortAllSets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 1; i <= 15; i++ {
			if _, err := NewString(strShortLen, NewStringOpt(i)); err != nil {
				b.Fatal("failed to generate pseudo-random string: ", err)
			}
		}
	}
}

func BenchmarkNewStrHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := NewString(strShortLen, WITH_LOWERCASE|WITH_DIGIT); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkNewStrHexUpper(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if _, err := NewString(strShortLen, WITH_UPPERCASE|WITH_DIGIT); err != nil {
			b.Fatal(err)
		}
	}
}
