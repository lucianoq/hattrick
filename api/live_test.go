package api

import (
	"reflect"
	"testing"

	"github.com/lucianoq/hattrick/chpp"
)

func Test_liveValues(t *testing.T) {
	tests := []struct {
		name                  string
		sourceSystem          chpp.SourceSystem
		includeStartingLineup bool
		useLiveEventsAndTexts bool
		want                  map[string]string
	}{
		{
			name: "all defaults omit their keys",
			want: map[string]string{},
		},
		{
			name:         "sourceSystem set",
			sourceSystem: chpp.SourceSystemYouthSystem,
			want:         map[string]string{"sourceSystem": "youth"},
		},
		{
			name:                  "both bools set",
			includeStartingLineup: true,
			useLiveEventsAndTexts: true,
			want: map[string]string{
				"includeStartingLineup": "true",
				"useLiveEventsAndTexts": "true",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := liveValues(tt.sourceSystem, tt.includeStartingLineup, tt.useLiveEventsAndTexts)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("liveValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
