package utils

import (
	"testing"
)

func TestGetwGoByString(t *testing.T) {

	str1 := "Hello, World!"

	counts1 := GetwGo(str1)

	if counts1["vowels"] != 3 {
		t.Errorf("Expected 3 vowels, got %d", counts1["vowels"])
	}
	if counts1["capital"] != 2 {
		t.Errorf("Expected 2 capital letters, got %d", counts1["capital"])
	}
	if counts1["small"] != 8 {
		t.Errorf("Expected 8 small letters, got %d", counts1["small"])
	}
	if counts1["spaces"] != 1 {
		t.Errorf("Expected 1 space, got %d", counts1["spaces"])
	}

}

func TestGetwGoBySpace(t *testing.T) {

	str2 := "   "

	counts2 := GetwGo(str2)

	if counts2["vowels"] != 0 {
		t.Errorf("Expected 0 vowels, got %d", counts2["vowels"])
	}
	if counts2["capital"] != 0 {
		t.Errorf("Expected 0 capital letters, got %d", counts2["capital"])
	}
	if counts2["small"] != 0 {
		t.Errorf("Expected 0 small letters, got %d", counts2["small"])
	}
	if counts2["spaces"] != 3 {
		t.Errorf("Expected 3 spaces, got %d", counts2["spaces"])
	}
}

func TestGetWAllSmallData(t *testing.T) {

	str2 := "thisisali"

	counts := GetwGo(str2)

	if counts["vowels"] != 4 {
		t.Errorf("Expected 4 vowels, got %d", counts["vowels"])
	}
	if counts["capital"] != 0 {
		t.Errorf("Expected 0 capital letters, got %d", counts["capital"])
	}
	if counts["small"] != 9 {
		t.Errorf("Expected 9 small letters, got %d", counts["small"])
	}
	if counts["spaces"] != 0 {
		t.Errorf("Expected 0 space, got %d", counts["spaces"])
	}

}

func TestGetWAllCapitalData(t *testing.T) {

	str2 := "THISISALI"

	counts := GetwGo(str2)

	if counts["vowels"] != 4 {
		t.Errorf("Expected 4 vowels, got %d", counts["vowels"])
	}
	if counts["capital"] != 9 {
		t.Errorf("Expected 9 capital letters, got %d", counts["capital"])
	}
	if counts["small"] != 0 {
		t.Errorf("Expected 0 small letters, got %d", counts["small"])
	}
	if counts["spaces"] != 0 {
		t.Errorf("Expected 0 space, got %d", counts["spaces"])
	}

}
