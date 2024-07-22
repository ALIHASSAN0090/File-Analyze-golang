package controllers

import (
	"fmt"
	"reflect"
	"testing"

	"main.go/utils"
)

func TestStats(t *testing.T) {

	fileContent := "Hello My Friend."

	expected := map[string]int{
		"vowels":  4,
		"capital": 3,
		"small":   10,
		"spaces":  2,
	}

	results := utils.GetwGo(fileContent) // it returns a map

	if !reflect.DeepEqual(expected, results) {
		t.Errorf("Expected %v but got %v", expected, results)
	} else {
		fmt.Printf("Expected %v and got %v", expected, results)
	}
}
