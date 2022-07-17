package string

import (
	"testing"
)

func Test_stringSize(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "string size",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stringSize()
		})
	}
}

func Test_traverseString(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "traverse string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			traverseString()
		})
	}
}

func Test_splitString(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "split string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			splitString()
		})
	}
}

func Test_createSlice(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "3 way create slice",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createSlice()
		})
	}
}

func Test_appendSlice(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "append slice",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			appendSlice()
		})
	}
}
