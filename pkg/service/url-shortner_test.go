package service

import "testing"

func TestKeyGenerator(t *testing.T) {

	tests := []struct {
		name  string
		arg   string
		want  string
		error bool
	}{
		{
			name:  "Success",
			arg:   "https://www.google.com",
			want:  "bPlNFG",
			error: false,
		},
		{
			name:  "Failure",
			arg:   "https://www.google.com",
			want:  "bPlNFG",
			error: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeyGenerator(); (got != tt.want) && (tt.error == false) {
				t.Errorf("KeyGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}
