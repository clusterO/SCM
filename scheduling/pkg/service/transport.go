package scheduling

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"
)

// Request and Response structs for scheduling service

type InsertSchedulingLogsRequest struct {
	Logs []SchedulingLog
}

type InsertSchedulingLogsResponse struct {
	Err error
}

type InsertRotationChangeRequest struct {
	RotationChange RotationChange
}

type InsertRotationChangeResponse struct {
	Err error
}

type InsertScheduleRequest struct {
	Schedule Schedule
}

type InsertScheduleResponse struct {
	ScheduleID int
	Err       error
}

type RemoveScheduleRequest struct {
	ScheduleID int
}

type RemoveScheduleResponse struct {
	Err error
}

type InsertShiftHistoryRequest struct {
	ShiftHistory ShiftHistory
}

type InsertShiftHistoryResponse struct {
	Err error
}

type UpdateScheduleRequest struct {
	Schedule Schedule
}

type UpdateScheduleResponse struct {
	Err error
}

type UpdateEmployeeShiftRequest struct {
	EmployeeShift EmployeeShift
}

type UpdateEmployeeShiftResponse struct {
	Err error
}

type StuckShiftRequest struct {
	ScheduleID int
}

type StuckShiftResponse struct {
	Err error
}

type AssignEmployeeToUrgentUnfilledRequest struct {
	ScheduleID int
	EmployeeID int
}

type AssignEmployeeToUrgentUnfilledResponse struct {
	Err error
}

type AddNoteScheduleRequest struct {
	ScheduleID int
	Note       string
}

type AddNoteScheduleResponse struct {
	Err error
}

type AddNoteMasterScheduleRequest struct {
	MasterScheduleID int
	Note             string
}

type AddNoteMasterScheduleResponse struct {
	Err error
}

type AddNoteMasterChangeRequest struct {
	MasterChangeID int
	Note           string
}

type AddNoteMasterChangeResponse struct {
	Err error
}

type GetMasterScheduleShiftsRequest struct {
	MasterScheduleID int
}

type GetMasterScheduleShiftsResponse struct {
	Shifts []Shift
	Err    error
}

type InsertDetailCallOutRequest struct {
	DetailCallOut DetailCallOut
}

type InsertDetailCallOutResponse struct {
	Err error
}

type InsertCallOutResultsRequest struct {
	CallOutResults []CallOutResult
}

type InsertCallOutResultsResponse struct {
	Err error
}

type AssignUserToLocationScheduleRequest struct {
	UserID     int
	LocationID int
}

type AssignUserToLocationScheduleResponse struct {
	Err error
}

type GetShiftsRequest struct {
	StartDate time.Time
	EndDate   time.Time
}

type GetShiftsResponse struct {
	Shifts []Shift
	Err    error
}

type PayPeriodArrayRequest struct {
	StartDate time.Time
	EndDate   time.Time
}

type PayPeriodArrayResponse struct {
	PayPeriod []time.Time
	Err       error
}

// Endpoints for scheduling service

type Endpoints struct {
	InsertSchedulingLogsEndpoint                 endpoint.Endpoint
	InsertRotationChangeEndpoint                 endpoint.Endpoint
	InsertScheduleEndpoint                       endpoint.Endpoint
	RemoveScheduleEndpoint                       endpoint.Endpoint
	InsertShiftHistoryEndpoint                   endpoint.Endpoint
	UpdateScheduleEndpoint                       endpoint.Endpoint
	UpdateEmployeeShiftEndpoint                  endpoint.Endpoint
	StuckShiftEndpoint                           endpoint.Endpoint
	AssignEmployeeToUrgentUnfilledEndpoint       endpoint.Endpoint
	AddNoteScheduleEndpoint                      endpoint.Endpoint
	AddNoteMasterScheduleEndpoint                endpoint.Endpoint
	AddNoteMasterChangeEndpoint                  endpoint.Endpoint
	GetMasterScheduleShiftsEndpoint              endpoint.Endpoint
	InsertDetailCallOutEndpoint                  endpoint.Endpoint
	InsertCallOutResultsEndpoint                 endpoint.Endpoint
	AssignUserToLocationScheduleEndpoint         endpoint.Endpoint
	GetShiftsEndpoint                            endpoint.Endpoint
	PayPeriodArrayEndpoint                       endpoint.Endpoint
}

// MakeEndpoints initializes all endpoints for the scheduling service
func MakeEndpoints(svc SchedulingService) Endpoints {
	return Endpoints{
		InsertSchedulingLogsEndpoint:                 makeInsertSchedulingLogsEndpoint(svc),
		InsertRotationChangeEndpoint:                 makeInsertRotationChangeEndpoint(svc),
		InsertScheduleEndpoint:                       makeInsertScheduleEndpoint(svc),
		RemoveScheduleEndpoint:                       makeRemoveScheduleEndpoint(svc),
		InsertShiftHistoryEndpoint:                   makeInsertShiftHistoryEndpoint(svc),
		UpdateScheduleEndpoint:                       makeUpdateScheduleEndpoint(svc),
		UpdateEmployeeShiftEndpoint:                  makeUpdateEmployeeShiftEndpoint(svc),
		StuckShiftEndpoint:                           makeStuckShiftEndpoint(svc),
		AssignEmployeeToUrgentUnfilledEndpoint:       makeAssignEmployeeToUrgentUnfilledEndpoint(svc),
		AddNoteScheduleEndpoint:                      makeAddNoteScheduleEndpoint(svc),
		AddNoteMasterScheduleEndpoint:                makeAddNoteMasterScheduleEndpoint(svc),
		AddNoteMasterChangeEndpoint:                  makeAddNoteMasterChangeEndpoint(svc),
		GetMasterScheduleShiftsEndpoint:              makeGetMasterScheduleShiftsEndpoint(svc),
		InsertDetailCallOutEndpoint:                  makeInsertDetailCallOutEndpoint(svc),
		InsertCallOutResultsEndpoint:                 makeInsertCallOutResultsEndpoint(svc),
		AssignUserToLocationScheduleEndpoint:         makeAssignUserToLocationScheduleEndpoint(svc),
		GetShiftsEndpoint:                            makeGetShiftsEndpoint(svc),
		PayPeriodArrayEndpoint:                       makePayPeriodArrayEndpoint(svc),
	}
}

// Endpoint functions for scheduling service 

func makeInsertSchedulingLogsEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InsertSchedulingLogsRequest)
		err := svc.InsertSchedulingLogs(ctx, req.Logs)
		return InsertSchedulingLogsResponse{Err: err}, nil
	}
}

func makeInsertRotationChangeEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InsertRotationChangeRequest)
		err := svc.InsertRotationChange(ctx, req.RotationChange)
		return InsertRotationChangeResponse{Err: err}, nil
	}
}

func makeInsertScheduleEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InsertScheduleRequest)
		scheduleID, err := svc.InsertSchedule(ctx, req.Schedule)
		return InsertScheduleResponse{ScheduleID: scheduleID, Err: err}, nil
	}
}

func makeRemoveScheduleEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RemoveScheduleRequest)
		err := svc.RemoveSchedule(ctx, req.ScheduleID)
		return RemoveScheduleResponse{Err: err}, nil
	}
}

func makeInsertShiftHistoryEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InsertShiftHistoryRequest)
		err := svc.InsertShiftHistory(ctx, req.ShiftHistory)
		return InsertShiftHistoryResponse{Err: err}, nil
	}
}

func makeUpdateScheduleEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateScheduleRequest)
		err := svc.UpdateSchedule(ctx, req.Schedule)
		return UpdateScheduleResponse{Err: err}, nil
	}
}

func makeUpdateEmployeeShiftEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateEmployeeShiftRequest)
		err := svc.UpdateEmployeeShift(ctx, req.EmployeeShift)
		return UpdateEmployeeShiftResponse{Err: err}, nil
	}
}

func makeStuckShiftEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StuckShiftRequest)
		err := svc.StuckShift(ctx, req.ScheduleID)
		return StuckShiftResponse{Err: err}, nil
	}
}

func makeAssignEmployeeToUrgentUnfilledEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AssignEmployeeToUrgentUnfilledRequest)
		err := svc.AssignEmployeeToUrgentUnfilled(ctx, req.ScheduleID, req.EmployeeID)
		return AssignEmployeeToUrgentUnfilledResponse{Err: err}, nil
	}
}

func makeAddNoteScheduleEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddNoteScheduleRequest)
		err := svc.AddNoteSchedule(ctx, req.ScheduleID, req.Note)
		return AddNoteScheduleResponse{Err: err}, nil
	}
}

func makeAddNoteMasterScheduleEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddNoteMasterScheduleRequest)
		err := svc.AddNoteMasterSchedule(ctx, req.MasterScheduleID, req.Note)
		return AddNoteMasterScheduleResponse{Err: err}, nil
	}
}

func makeAddNoteMasterChangeEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddNoteMasterChangeRequest)
		err := svc.AddNoteMasterChange(ctx, req.MasterChangeID, req.Note)
		return AddNoteMasterChangeResponse{Err: err}, nil
	}
}

func makeGetMasterScheduleShiftsEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetMasterScheduleShiftsRequest)
		shifts, err := svc.GetMasterScheduleShifts(ctx, req.MasterScheduleID)
		return GetMasterScheduleShiftsResponse{Shifts: shifts, Err: err}, nil
	}
}

func makeInsertDetailCallOutEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InsertDetailCallOutRequest)
		err := svc.InsertDetailCallOut(ctx, req.DetailCallOut)
		return InsertDetailCallOutResponse{Err: err}, nil
	}
}

func makeInsertCallOutResultsEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(InsertCallOutResultsRequest)
		err := svc.InsertCallOutResults(ctx, req.CallOutResults)
		return InsertCallOutResultsResponse{Err: err}, nil
	}
}

func makeAssignUserToLocationScheduleEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AssignUserToLocationScheduleRequest)
		err := svc.AssignUserToLocationSchedule(ctx, req.LocationID, req.UserID)
		return AssignUserToLocationScheduleResponse{Err: err}, nil
	}
}

func makeGetShiftsEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetShiftsRequest)
		shifts, err := svc.GetShifts(ctx, req.ScheduleID, req.StartTime, req.EndTime)
		return GetShiftsResponse{Shifts: shifts, Err: err}, nil
	}
}

func makePayPeriodArrayEndpoint(svc SchedulingService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PayPeriodArrayRequest)
		payPeriods, err := svc.PayPeriodArray(ctx, req.StartTime, req.EndTime)
		return PayPeriodArrayResponse{PayPeriods: payPeriods, Err: err}, nil
	}
}

func decodeInsertSchedulingLogsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req InsertSchedulingLogsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeInsertRotationChangeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req InsertRotationChangeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeInsertScheduleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req InsertScheduleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeRemoveScheduleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req RemoveScheduleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeInsertShiftHistoryRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req InsertShiftHistoryRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUpdateScheduleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req UpdateScheduleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeUpdateEmployeeShiftRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req UpdateEmployeeShiftRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeStuckShiftRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req StuckShiftRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeAssignEmployeeToUrgentUnfilledRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req AssignEmployeeToUrgentUnfilledRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeAddNoteScheduleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req AddNoteScheduleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeAddNoteMasterScheduleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req AddNoteMasterScheduleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeAddNoteMasterChangeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req AddNoteMasterChangeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Decoder function for GetMasterScheduleShiftsRequest
func decodeGetMasterScheduleShiftsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetMasterScheduleShiftsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Decoder function for InsertDetailCallOutRequest
func decodeInsertDetailCallOutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req InsertDetailCallOutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Decoder function for InsertCallOutResultsRequest
func decodeInsertCallOutResultsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req InsertCallOutResultsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Decoder function for AssignUserToLocationScheduleRequest
func decodeAssignUserToLocationScheduleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req AssignUserToLocationScheduleRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Decoder function for GetShiftsRequest
func decodeGetShiftsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req GetShiftsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Decoder function for PayPeriodArrayRequest
func decodePayPeriodArrayRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req PayPeriodArrayRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Encoder function for responses
func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
	}