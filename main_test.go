package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_split(t *testing.T) {
	tests := []struct {
		name       string
		line       string
		substrings []string
		answer     []string
	}{
		{
			name:       "success",
			line:       "123456",
			substrings: []string{"123", "12", "34", "56", "45", "6", "4"},
			answer:     []string{"123", "45", "6"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := split(tt.line, tt.substrings)
			if !reflect.DeepEqual(out, tt.answer) {
				t.Errorf("[%s]\nwant: %v\n got: %v", tt.name, tt.answer, out)
			}
		})
	}
}

func Test_sort(t *testing.T) {

	tests := []struct {
		name   string
		line   string
		arr    []string
		answer []string
	}{
		{
			name:   "success",
			arr:    []string{"4", "6", "12", "34", "56", "45", "123"},
			answer: []string{"123", "12", "34", "56", "45", "6", "4"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortedArr := make([]string, len(tt.arr))
			copy(sortedArr, tt.arr)
			sort.Sort(byLen(sortedArr))

			if !reflect.DeepEqual(sortedArr, tt.answer) {
				t.Errorf("[%s]\nwant: %v\n got: %v", tt.name, tt.answer, sortedArr)
			}
		})
	}
}
