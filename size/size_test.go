package size

import "testing"

func Test_intSize(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "int size",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			intSize()
		})
	}
}

func Test_uintSize(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "uint size",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uintSize()
		})
	}
}

func Test_emptyStructSize(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "empty struct size",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			emptyStructSize()
		})
	}
}

func Test_normalStructSize(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "normal struct size",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalStructSize()
		})
	}
}

func Test_pointerSize(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "pointer size",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pointerSize()
		})
	}
}
