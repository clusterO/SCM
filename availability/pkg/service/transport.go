package availability

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type CheckAvailabilityRequest struct {
	Item string
}

type CheckAvailabilityResponse struct {
	Available bool
	Error     error
}

type GetAvailabilityRequest struct {
	Item string
}

type GetAvailabilityResponse struct {
	Count int
	Error error
}

type ReserveRequest struct {
	Item string
}

type ReserveResponse struct {
	Err error
}

func MakeGetAvailabilityEndpoint(svc AvailabilityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAvailabilityRequest)
		availability, err := svc.GetAvailability(ctx, req.Item)
		if err != nil {
			return GetAvailabilityResponse{Count: 0, Error: err}, nil
		}
		return GetAvailabilityResponse{Count: availability, Error: err}, nil
	}
}

func MakeCheckAvailabilityEndpoint(svc AvailabilityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CheckAvailabilityRequest)
		available, err := svc.CheckAvailability(ctx, req.Item)
		if err != nil {
			return CheckAvailabilityResponse{Available: false, Error: err}, nil
		}
		return CheckAvailabilityResponse{Available: available, Error: err}, nil
	}
}

func MakeReserveAvailabilityEndpoint(svc AvailabilityService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReserveRequest)
		err := svc.Reserve(ctx, req.Item)
		if err != nil {
			return ReserveResponse{Err: err}, nil
		}
		return ReserveResponse{Err: err}, nil
	}
}
