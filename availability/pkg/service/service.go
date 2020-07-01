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
	// Implement your logic to check the availability of the item here
	// For this example, let's assume the item is always available
	return true, nil
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
