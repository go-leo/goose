package goose

import "testing"

// TestFormatBool is a test function for FormatBool.
func TestFormatBool(t *testing.T) {
	tests := []struct {
		name string
		b    bool
		want string
	}{
		{"true", true, "true"},
		{"false", false, "false"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatBool(tt.b); got != tt.want {
				t.Errorf("FormatBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 下面是单元测试代码
func TestFormatBoolSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []bool // 使用原生bool类型，泛型参数会在测试中实例化
		expected []string
	}{
		{
			name:     "nil slice",
			input:    nil,
			expected: nil,
		},
		{
			name:     "empty slice",
			input:    []bool{},
			expected: []string{},
		},
		{
			name:     "slice with true",
			input:    []bool{true},
			expected: []string{"true"},
		},
		{
			name:     "slice with false",
			input:    []bool{false},
			expected: []string{"false"},
		},
		{
			name:     "slice with true and false",
			input:    []bool{true, false},
			expected: []string{"true", "false"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := FormatBoolSlice(tt.input) // 使用测试数据调用待测函数
			if len(actual) != len(tt.expected) {
				t.Errorf("FormatBoolSlice(%v) expected length %d, actual length %d", tt.input, len(tt.expected), len(actual))
			}
			for i, v := range actual {
				if v != tt.expected[i] {
					t.Errorf("FormatBoolSlice(%v) expected %v, actual %v at index %d", tt.input, tt.expected, actual, i)
				}
			}
		})
	}
}
