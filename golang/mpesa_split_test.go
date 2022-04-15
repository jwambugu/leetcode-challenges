package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMpesaSplit(t *testing.T) {
	tests := []struct {
		amount         float64
		wantsSchedules []schedule
		wantsBalance   float64
	}{
		{300, []schedule{{300, 0}}, 0},
		{150_000, []schedule{{150_000, 0}}, 0},
		{300_000, []schedule{{150_000, 0}, {150_000, 0}}, 0},
		{300_099, []schedule{
			{150_000, 0}, {150_000, 0},
		}, 99},
		{500_000, []schedule{
			{150_000, 0}, {150_000, 0},
			{150_000, 1}, {50_000, 1},
		}, 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("splitting %d", int64(tt.amount)), func(t *testing.T) {
			balance, schedules := mpesaSplit([]schedule{}, tt.amount, 0, false)

			if balance != tt.wantsBalance {
				t.Fatalf("got balance %f, want %f", balance, tt.wantsBalance)
			}

			if !reflect.DeepEqual(schedules, tt.wantsSchedules) {
				t.Fatalf("got schedules %v, want %v", schedules, tt.wantsSchedules)
			}
		})
	}
}
