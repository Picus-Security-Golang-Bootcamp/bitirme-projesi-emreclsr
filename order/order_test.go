package order_test

import (
	"github.com/emreclsr/picusfinal/order"
	"github.com/go-playground/assert/v2"

	"testing"
	"time"
)

func TestCheckTime(t *testing.T) {
	tests := []struct {
		name string
		day  int
		want bool
	}{
		{name: "test1", day: 1, want: true},
		{name: "test5", day: 5, want: true},
		{name: "test10", day: 10, want: true},
		{name: "test15", day: 15, want: false},
		{name: "test20", day: 20, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var order order.Order
			timeDuration := time.Duration(tt.day) * time.Hour * 24
			order.CreatedAt = time.Now().Add(-timeDuration)
			assert.Equal(t, tt.want, order.CheckTime())
		})
	}
}
