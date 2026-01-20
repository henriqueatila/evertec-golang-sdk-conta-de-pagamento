package client

import "time"

// Test helper functions for pointer types

func strPtr(s string) *string {
	return &s
}

func intPtr(i int64) *int64 {
	return &i
}

func timePtr(t time.Time) *time.Time {
	return &t
}
