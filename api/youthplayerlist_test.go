package api

import (
	"reflect"
	"testing"
)

func Test_youthPlayerDetailsValues(t *testing.T) {
	tests := []struct {
		name          string
		showScoutCall bool
		showLastMatch bool
		want          map[string]string
	}{
		{
			name: "both false omit their keys",
			want: map[string]string{},
		},
		{
			name:          "showScoutCall set",
			showScoutCall: true,
			want:          map[string]string{"showScoutCall": "true"},
		},
		{
			name:          "both set",
			showScoutCall: true,
			showLastMatch: true,
			want: map[string]string{
				"showScoutCall": "true",
				"showLastMatch": "true",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := youthPlayerDetailsValues(tt.showScoutCall, tt.showLastMatch)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("youthPlayerDetailsValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
