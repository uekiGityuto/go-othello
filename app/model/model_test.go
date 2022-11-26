package model

import (
	"fmt"
	"testing"
)

func TestColor_Valid(t *testing.T) {
	type expected struct {
		err error
	}

	tests := []struct {
		name     string
		input    Color
		expected expected
	}{
		{
			name:  "正常時: 黒がvalidになること",
			input: Black,
			expected: expected{
				err: nil,
			},
		},
		{
			name:  "正常時: 白がvalidになること",
			input: White,
			expected: expected{
				err: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.input.Valid(); err != nil {
				t.Errorf("want no error, but got %v", err)
			}
		})
	}

	errorTests := []struct {
		name     string
		input    Color
		expected expected
	}{
		{
			name:  "異常時: 不適切な数字の場合にerrorを返却すること",
			input: 9,
			expected: expected{
				err: fmt.Errorf(`"%d" is invalid as color`, 9),
			},
		},
	}

	for _, tt := range errorTests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.input.Valid()
			if err == nil {
				t.Errorf("want error, but no error")
			}
		})
	}
}
