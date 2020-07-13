package availability

import (
	"context"
	"errors"
)

type AvailabilityService interface {
	CheckAvailability(context.Context, string) (bool, error)
	GetAvailability(context.Context, string) (int, error)
	Reserve(context.Context, string) error
}

type availabilityService struct{}

func (availabilityService) CheckAvailability(ctx context.Context, item string) (bool, error) {
	// For demonstration purposes, let's assume the item is always available
	// You should replace this with your actual implementation

	// If the item is available, return true
	// If the item is not available, return false
	// If there's an error while checking availability, return an error

	// Example implementation:
	if item == "example_item" {
		return true, nil
	}

	// If the item is not available, return false
	return false, nil
}

func (availabilityService) GetAvailability(ctx context.Context, item string) (int, error) {
	// Implement your logic to get the availability count of the item here
	// For this example, let's assume the availability count is 10
	return 10, nil
}

func (availabilityService) Reserve(ctx context.Context, item string) error {
	// Implement your logic to reserve the item here
	// For this example, let's assume the reservation is successful
	return nil
}

// ErrItemNotFound is returned when the item is not found
var ErrItemNotFound = errors.New("Item not found")
