package api

import (
	"reflect"
	"testing"
	"time"
)

func Test_createBins(t *testing.T) {

	var (
		today     = time.Now()
		daysAgo49 = today.Add(-49 * 24 * time.Hour)
		daysAgo70 = today.Add(-70 * 24 * time.Hour)
		yearAgo2  = today.Add(-2 * 365 * 24 * time.Hour)
	)

	t.Parallel()
	tests := []struct {
		name  string
		start time.Time
		end   time.Time
		want  []interval
	}{
		{
			name:  "less than 50 days, no split",
			start: daysAgo49,
			end:   today,
			want: []interval{
				{daysAgo49, today},
			},
		},
		{
			name:  "70 days ago -> today, split 2",
			start: daysAgo70,
			end:   today,
			want: []interval{
				{daysAgo70, daysAgo70.Add(binMaxSize)},
				{daysAgo70.Add(binMaxSize).Add(time.Nanosecond), today},
			},
		},
		{
			name:  "2 year ago -> today, split many",
			start: yearAgo2,
			end:   today,
			want: []interval{
				{
					from: yearAgo2,
					to:   yearAgo2.Add(binMaxSize),
				},
				{
					from: yearAgo2.Add(binMaxSize).Add(time.Nanosecond),
					to:   yearAgo2.Add(2 * binMaxSize),
				},
				{
					from: yearAgo2.Add(2 * binMaxSize).Add(time.Nanosecond),
					to:   yearAgo2.Add(3 * binMaxSize),
				},
				{
					from: yearAgo2.Add(3 * binMaxSize).Add(time.Nanosecond),
					to:   yearAgo2.Add(4 * binMaxSize),
				},
				{
					from: yearAgo2.Add(4 * binMaxSize).Add(time.Nanosecond),
					to:   yearAgo2.Add(5 * binMaxSize),
				},
				{
					from: yearAgo2.Add(5 * binMaxSize).Add(time.Nanosecond),
					to:   yearAgo2.Add(6 * binMaxSize),
				},
				{
					from: yearAgo2.Add(6 * binMaxSize).Add(time.Nanosecond),
					to:   yearAgo2.Add(7 * binMaxSize),
				},
				{
					from: yearAgo2.Add(7 * binMaxSize).Add(time.Nanosecond),
					to:   yearAgo2.Add(8 * binMaxSize),
				},
				{
					from: yearAgo2.Add(8 * binMaxSize).Add(time.Nanosecond),
					to:   yearAgo2.Add(9 * binMaxSize),
				},
				{
					from: yearAgo2.Add(9 * binMaxSize).Add(time.Nanosecond),
					to:   yearAgo2.Add(10 * binMaxSize),
				},
				{
					from: yearAgo2.Add(10 * binMaxSize).Add(time.Nanosecond),
					to:   yearAgo2.Add(11 * binMaxSize),
				},
				{
					from: yearAgo2.Add(11 * binMaxSize).Add(time.Nanosecond),
					to:   yearAgo2.Add(12 * binMaxSize),
				},
				{
					from: yearAgo2.Add(12 * binMaxSize).Add(time.Nanosecond),
					to:   yearAgo2.Add(13 * binMaxSize),
				},
				{
					from: yearAgo2.Add(13 * binMaxSize).Add(time.Nanosecond),
					to:   yearAgo2.Add(14 * binMaxSize),
				},
				{
					from: yearAgo2.Add(14 * binMaxSize).Add(time.Nanosecond),
					to:   today,
				},
			},
		},
	}
	for _, tt := range tests {
		var tt = tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := createBins(tt.start, tt.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createBins() = %v, want %v", got, tt.want)
			}
		})
	}
}
