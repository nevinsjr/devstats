package utils

import (
	"testing"
	"time"
)

func TestActivityIsGreaterThanOneWeek(t *testing.T) {
	// Arrange
	duration, _ := time.ParseDuration("192h")

	// Act
	isGreater := IsActivityGreaterThanOneWeek(duration)

	// Assert
	if !isGreater {
		t.Errorf("Expected %v but got %v", true, isGreater)
	}
}
