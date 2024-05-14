package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountBeefOccurrences(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int32
	}{
		{"Empty string", "", map[string]int32{}},
		{"Same word differnt case", "Ribeye ribeye", map[string]int32{"ribeye": 2}},
		{"Words separate by comma", "ribeye,ribeye", map[string]int32{"ribeye": 2}},
		{"Words separate by dot", "ribeye.ribeye", map[string]int32{"ribeye": 2}},
		{"Words separate by multiple white spaces", " ribeye	\r \nribeye ", map[string]int32{"ribeye": 2}},
		{"With Hyphen", "T-bone t-bone", map[string]int32{"t-bone": 2}},
		{"Multiple beefs", "Tenderloin ribeye T-bone t-bone", map[string]int32{"tenderloin": 1, "ribeye": 1, "t-bone": 2}},
		{"Mix", " Tenderloin ,	.ribeye T-bone. .,t-bone .", map[string]int32{"tenderloin": 1, "ribeye": 1, "t-bone": 2}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := countBeefOccurrences(test.input)
			if !reflect.DeepEqual(actual, test.expected) {
				expectedMapStr := fmt.Sprintf("%v", test.expected)
				actualMapStr := fmt.Sprintf("%v", actual)
				t.Errorf("Test '%s' failed: Expected %s, got %s", test.name, expectedMapStr, actualMapStr)
			}
		})
	}
}
