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
	// For demonstration purposes, let's assume the availability count is always 10

	// If the availability count is obtained successfully, return it
	// If there's an error while fetching the availability count, return an error

	// Example implementation:
	if item == "example_item" {
		return 10, nil
	}

	// If the availability count cannot be obtained, return an error
	return 0, errors.New("failed to get availability count")
}

func (availabilityService) Reserve(ctx context.Context, item string) error {
	// For demonstration purposes, let's assume the reservation is successful

	// If the item is successfully reserved, return nil
	// If there's an error while reserving the item, return an error

	// Example implementation:
	if item == "example_item" {
		return nil
	}

	// If the item cannot be reserved, return an error
	return errors.New("failed to reserve item")
}

// ErrItemNotFound is returned when the item is not found
var ErrItemNotFound = errors.New("Item not found")
