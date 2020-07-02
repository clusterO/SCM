package availability

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

