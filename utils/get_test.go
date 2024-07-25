package utils

import (
	"sync"
	"testing"
)

// function testing by passing string
func TestGetStringData(t *testing.T) {

	wg := &sync.WaitGroup{}
	results := make(chan map[string]int)

	go GetData("Hello, World!", wg, results)

	wg.Add(1)

	go func() {
		wg.Wait()
		close(results)
	}()

	for counts := range results {
		if counts["vowels"] != 3 {
			t.Errorf("Expected 3 vowels, got %d", counts["vowels"])
		}
		if counts["capital"] != 2 {
			t.Errorf("Expected 2 capital letters, got %d", counts["capital"])
		}
		if counts["small"] != 8 {
			t.Errorf("Expected 8 small letters, got %d", counts["small"])
		}
		if counts["spaces"] != 1 {
			t.Errorf("Expected 1 space, got %d", counts["spaces"])
		}
	}
}

//function testing by passing spaces.

func TestGetSpaceData(t *testing.T) {
	wg := &sync.WaitGroup{}
	results := make(chan map[string]int)
	go GetData("   ", wg, results)
	wg.Add(1)
	go func() {
		wg.Wait()
		close(results)
	}()

	for counts := range results {
		if counts["vowels"] != 0 {
			t.Errorf("Expected 0 vowels, got %d", counts["vowels"])
		}
		if counts["capital"] != 0 {
			t.Errorf("Expected 0 capital letters, got %d", counts["capital"])
		}
		if counts["small"] != 0 {
			t.Errorf("Expected 0 small letters, got %d", counts["small"])
		}
		if counts["spaces"] != 3 {
			t.Errorf("Expected 3 space, got %d", counts["spaces"])
		}
	}
}

func TestGetAllSmallData(t *testing.T) {
	wg := &sync.WaitGroup{}
	results := make(chan map[string]int)
	go GetData("thisisali", wg, results)
	wg.Add(1)
	go func() {
		wg.Wait()
		close(results)
	}()

	for counts := range results {
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
}

func TestGetAllCapitalData(t *testing.T) {
	wg := &sync.WaitGroup{}
	results := make(chan map[string]int)
	go GetData("THISISALI", wg, results)
	wg.Add(1)
	go func() {
		wg.Wait()
		close(results)
	}()

	for counts := range results {
		if counts["vowels"] != 4 {
			t.Errorf("Expected 4 vowels, got %d", counts["vowels"])
		}
		if counts["capital"] != 9 {
			t.Errorf("Expected 0 capital letters, got %d", counts["capital"])
		}
		if counts["small"] != 0 {
			t.Errorf("Expected 0 small letters, got %d", counts["small"])
		}
		if counts["spaces"] != 0 {
			t.Errorf("Expected 0 space, got %d", counts["spaces"])
		}
	}
}
