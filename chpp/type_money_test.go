package chpp

import "testing"

func TestMoney_String(t *testing.T) {
	tests := []struct {
		m    Money
		want string
	}{
		{0, "0 kr"},
		{999, "999 kr"},
		{1000, "1,000 kr"},
		{1234567, "1,234,567 kr"},
		{-1500, "-1,500 kr"},
	}

	for _, tt := range tests {
		if got := tt.m.String(); got != tt.want {
			t.Errorf("Money(%d).String() = %q, want %q", tt.m, got, tt.want)
		}
	}
}
