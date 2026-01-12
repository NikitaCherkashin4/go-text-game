package models

import (
	"strings"
	"testing"
)

func TestGetPlayerName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "user says yes and provides name",
			input: "yes\nAlice\n",
			want:  "Alice",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockInput := strings.NewReader(tt.input)
			got := GetPlayerName(mockInput)
			if got != tt.want {
				t.Errorf("GetPlayerName() = %v, want %v", got, tt.want)
			}
		})
	}
}
