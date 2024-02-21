package errors

import (
	"errors"
	"sm_errors/entities"
	entities_grpc "sm_errors/entities/grpc"
	entities_http "sm_errors/entities/http"
	entities_ws "sm_errors/entities/ws"
	"testing"
)

var (
	ExampleError = Constructor{
		ID:      1000000000000000,
		Status:  entities.StatusFailed,
		Message: new(entities.Message).Set("Example error message. "),

		Grpc: GrpcConstructor{
			StatusCode: entities_grpc.StatusInternalServerError,
		},
		Http: HttpConstructor{
			StatusCode: entities_http.StatusInternalServerError,
		},
		Ws: WsConstructor{
			StatusCode: entities_ws.StatusInternalServerError,
		},
	}.Build()
	ExampleErrorEmptyMessage = Constructor{
		ID:      1000000000000001,
		Status:  entities.StatusFailed,
		Message: new(entities.Message),

		Grpc: GrpcConstructor{
			StatusCode: entities_grpc.StatusInternalServerError,
		},
		Http: HttpConstructor{
			StatusCode: entities_http.StatusInternalServerError,
		},
		Ws: WsConstructor{
			StatusCode: entities_ws.StatusInternalServerError,
		},
	}.Build()

	ExampleFieldsError = FieldsConstructor{
		ID:      1000000000000000,
		Status:  entities.StatusFailed,
		Message: new(entities.Message).Set("Example fields error message. "),

		Grpc: GrpcConstructor{
			StatusCode: entities_grpc.StatusInternalServerError,
		},
		Http: HttpConstructor{
			StatusCode: entities_http.StatusInternalServerError,
		},
		Ws: WsConstructor{
			StatusCode: entities_ws.StatusInternalServerError,
		},
	}.Build()
	ExampleFieldsErrorEmptyMessage = FieldsConstructor{
		ID:      1000000000000001,
		Status:  entities.StatusFailed,
		Message: new(entities.Message),

		Grpc: GrpcConstructor{
			StatusCode: entities_grpc.StatusInternalServerError,
		},
		Http: HttpConstructor{
			StatusCode: entities_http.StatusInternalServerError,
		},
		Ws: WsConstructor{
			StatusCode: entities_ws.StatusInternalServerError,
		},
	}.Build()
)

func TestError(t *testing.T) {
	var (
		srcErr = errors.New("Test error. ")
		err    = ExampleError.SetError(srcErr)
	)

	if err.ID() != 1000000000000000 {
		t.Fatal("Invalid error id. ")
	}

	if err.String() != "Example error message. " {
		t.Fatal("Invalid string error. ")
	}

	if err.Message() != "Example error message. " {
		t.Fatal("Invalid error message. ")
	}

	if !err.Is(srcErr) {
		t.Fatal("Invalid source error. ")
	}

	if err.Status() != entities.StatusFailed {
		t.Fatal("Invalid error status. ")
	}

	if err.GrpcStatusCode() != entities_grpc.StatusInternalServerError {
		t.Fatal("Invalid error grpc status code. ")
	}

	if err.HttpStatusCode() != entities_http.StatusInternalServerError {
		t.Fatal("Invalid error http status code. ")
	}

	if err.WsStatusCode() != entities_ws.StatusInternalServerError {
		t.Fatal("Invalid error ws status code. ")
	}

	err = ExampleErrorEmptyMessage.SetError(srcErr)

	if err.ID() != 1000000000000001 {
		t.Fatal("Invalid error id. ")
	}

	if err.String() != "Test error. " {
		t.Fatal("Invalid string error. ")
	}

	if err.Message() != "" {
		t.Fatal("Invalid error message. ")
	}
}

func TestGrpcError(t *testing.T) {
	var (
		srcErr = errors.New("Test error. ")
		err    = Grpc(ExampleError.SetError(srcErr))
	)

	if err.ID() != 1000000000000000 {
		t.Fatal("Invalid error id. ")
	}

	if err.String() != "Example error message. " {
		t.Fatal("Invalid string error. ")
	}

	if err.Message() != "Example error message. " {
		t.Fatal("Invalid error message. ")
	}

	if !err.Is(srcErr) {
		t.Fatal("Invalid source error. ")
	}

	if err.Status() != entities.StatusFailed {
		t.Fatal("Invalid error status. ")
	}

	if err.GrpcStatusCode() != entities_grpc.StatusInternalServerError {
		t.Fatal("Invalid error grpc status code. ")
	}

	err = Grpc(ExampleErrorEmptyMessage.SetError(srcErr))

	if err.ID() != 1000000000000001 {
		t.Fatal("Invalid error id. ")
	}

	if err.String() != "Test error. " {
		t.Fatal("Invalid string error. ")
	}

	if err.Message() != "" {
		t.Fatal("Invalid error message. ")
	}
}

func TestHttpError(t *testing.T) {
	var (
		srcErr = errors.New("Test error. ")
		err    = Http(ExampleError.SetError(srcErr))
	)

	if err.ID() != 1000000000000000 {
		t.Fatal("Invalid error id. ")
	}

	if err.String() != "Example error message. " {
		t.Fatal("Invalid string error. ")
	}

	if err.Message() != "Example error message. " {
		t.Fatal("Invalid error message. ")
	}

	if !err.Is(srcErr) {
		t.Fatal("Invalid source error. ")
	}

	if err.Status() != entities.StatusFailed {
		t.Fatal("Invalid error status. ")
	}

	if err.HttpStatusCode() != entities_http.StatusInternalServerError {
		t.Fatal("Invalid error http status code. ")
	}

	err = Http(ExampleErrorEmptyMessage.SetError(srcErr))

	if err.ID() != 1000000000000001 {
		t.Fatal("Invalid error id. ")
	}

	if err.String() != "Test error. " {
		t.Fatal("Invalid string error. ")
	}

	if err.Message() != "" {
		t.Fatal("Invalid error message. ")
	}
}

func TestWsError(t *testing.T) {
	var (
		srcErr = errors.New("Test error. ")
		err    = Ws(ExampleError.SetError(srcErr))
	)

	if err.ID() != 1000000000000000 {
		t.Fatal("Invalid error id. ")
	}

	if err.String() != "Example error message. " {
		t.Fatal("Invalid string error. ")
	}

	if err.Message() != "Example error message. " {
		t.Fatal("Invalid error message. ")
	}

	if !err.Is(srcErr) {
		t.Fatal("Invalid source error. ")
	}

	if err.Status() != entities.StatusFailed {
		t.Fatal("Invalid error status. ")
	}

	if err.WsStatusCode() != entities_ws.StatusInternalServerError {
		t.Fatal("Invalid error ws status code. ")
	}

	err = Ws(ExampleErrorEmptyMessage.SetError(srcErr))

	if err.ID() != 1000000000000001 {
		t.Fatal("Invalid error id. ")
	}

	if err.String() != "Test error. " {
		t.Fatal("Invalid string error. ")
	}

	if err.Message() != "" {
		t.Fatal("Invalid error message. ")
	}
}

func TestFieldsError(t *testing.T) {
	var (
		srcErr = errors.New("Test fields error. ")
		err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
			Key:     "test",
			Message: "Invalid value",
		})
	)

	if err.ID() != 1000000000000000 {
		t.Fatal("Invalid error id. ")
	}

	if err.String() != "Example fields error message. " {
		t.Fatal("Invalid string error. ")
	}

	if err.Message() != "Example fields error message. " {
		t.Fatal("Invalid error message. ")
	}

	if !err.Is(srcErr) {
		t.Fatal("Invalid source error. ")
	}

	if err.Status() != entities.StatusFailed {
		t.Fatal("Invalid error status. ")
	}

	if err.GrpcStatusCode() != entities_grpc.StatusInternalServerError {
		t.Fatal("Invalid error grpc status code. ")
	}

	if err.HttpStatusCode() != entities_http.StatusInternalServerError {
		t.Fatal("Invalid error http status code. ")
	}

	if err.WsStatusCode() != entities_ws.StatusInternalServerError {
		t.Fatal("Invalid error ws status code. ")
	}

	if len(err.Fields()) != 1 {
		t.Fatal("Invalid error fields. ")
	}

	if field := err.Fields().Get("test"); field.Key != "test" || field.Message != "Invalid value" {
		t.Fatal("Invalid error fields. ")
	}

	err = ExampleFieldsErrorEmptyMessage.SetError(srcErr)

	if err.ID() != 1000000000000001 {
		t.Fatal("Invalid error id. ")
	}

	if err.String() != "Test fields error. " {
		t.Fatal("Invalid string error. ")
	}

	if err.Message() != "" {
		t.Fatal("Invalid error message. ")
	}

	if len(err.Fields()) != 0 {
		t.Fatal("Invalid error fields. ")
	}

	if field := err.Fields().Get("test"); field.Key != "" || field.Message != "" {
		t.Fatal("Invalid error fields. ")
	}
}

func TestGrpcFieldsError(t *testing.T) {
	var (
		srcErr = errors.New("Test fields error. ")
		err    = FieldsGrpc(ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
			Key:     "test",
			Message: "Invalid value",
		}))
	)

	if err.ID() != 1000000000000000 {
		t.Fatal("Invalid error id. ")
	}

	if err.String() != "Example fields error message. " {
		t.Fatal("Invalid string error. ")
	}

	if err.Message() != "Example fields error message. " {
		t.Fatal("Invalid error message. ")
	}

	if !err.Is(srcErr) {
		t.Fatal("Invalid source error. ")
	}

	if err.Status() != entities.StatusFailed {
		t.Fatal("Invalid error status. ")
	}

	if err.GrpcStatusCode() != entities_grpc.StatusInternalServerError {
		t.Fatal("Invalid error grpc status code. ")
	}

	if len(err.Fields()) != 1 {
		t.Fatal("Invalid error fields. ")
	}

	if field := err.Fields().Get("test"); field.Key != "test" || field.Message != "Invalid value" {
		t.Fatal("Invalid error fields. ")
	}

	err = FieldsGrpc(ExampleFieldsErrorEmptyMessage.SetError(srcErr))

	if err.ID() != 1000000000000001 {
		t.Fatal("Invalid error id. ")
	}

	if err.String() != "Test fields error. " {
		t.Fatal("Invalid string error. ")
	}

	if err.Message() != "" {
		t.Fatal("Invalid error message. ")
	}

	if len(err.Fields()) != 0 {
		t.Fatal("Invalid error fields. ")
	}

	if field := err.Fields().Get("test"); field.Key != "" || field.Message != "" {
		t.Fatal("Invalid error fields. ")
	}
}

func TestHttpFieldsError(t *testing.T) {
	var (
		srcErr = errors.New("Test fields error. ")
		err    = FieldsHttp(ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
			Key:     "test",
			Message: "Invalid value",
		}))
	)

	if err.ID() != 1000000000000000 {
		t.Fatal("Invalid error id. ")
	}

	if err.String() != "Example fields error message. " {
		t.Fatal("Invalid string error. ")
	}

	if err.Message() != "Example fields error message. " {
		t.Fatal("Invalid error message. ")
	}

	if !err.Is(srcErr) {
		t.Fatal("Invalid source error. ")
	}

	if err.Status() != entities.StatusFailed {
		t.Fatal("Invalid error status. ")
	}

	if err.HttpStatusCode() != entities_http.StatusInternalServerError {
		t.Fatal("Invalid error http status code. ")
	}

	if len(err.Fields()) != 1 {
		t.Fatal("Invalid error fields. ")
	}

	if field := err.Fields().Get("test"); field.Key != "test" || field.Message != "Invalid value" {
		t.Fatal("Invalid error fields. ")
	}

	err = FieldsHttp(ExampleFieldsErrorEmptyMessage.SetError(srcErr))

	if err.ID() != 1000000000000001 {
		t.Fatal("Invalid error id. ")
	}

	if err.String() != "Test fields error. " {
		t.Fatal("Invalid string error. ")
	}

	if err.Message() != "" {
		t.Fatal("Invalid error message. ")
	}

	if len(err.Fields()) != 0 {
		t.Fatal("Invalid error fields. ")
	}

	if field := err.Fields().Get("test"); field.Key != "" || field.Message != "" {
		t.Fatal("Invalid error fields. ")
	}
}

func TestWsFieldsError(t *testing.T) {
	var (
		srcErr = errors.New("Test fields error. ")
		err    = FieldsWs(ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
			Key:     "test",
			Message: "Invalid value",
		}))
	)

	if err.ID() != 1000000000000000 {
		t.Fatal("Invalid error id. ")
	}

	if err.String() != "Example fields error message. " {
		t.Fatal("Invalid string error. ")
	}

	if err.Message() != "Example fields error message. " {
		t.Fatal("Invalid error message. ")
	}

	if !err.Is(srcErr) {
		t.Fatal("Invalid source error. ")
	}

	if err.Status() != entities.StatusFailed {
		t.Fatal("Invalid error status. ")
	}

	if err.WsStatusCode() != entities_ws.StatusInternalServerError {
		t.Fatal("Invalid error ws status code. ")
	}

	if len(err.Fields()) != 1 {
		t.Fatal("Invalid error fields. ")
	}

	if field := err.Fields().Get("test"); field.Key != "test" || field.Message != "Invalid value" {
		t.Fatal("Invalid error fields. ")
	}

	err = FieldsWs(ExampleFieldsErrorEmptyMessage.SetError(srcErr))

	if err.ID() != 1000000000000001 {
		t.Fatal("Invalid error id. ")
	}

	if err.String() != "Test fields error. " {
		t.Fatal("Invalid string error. ")
	}

	if err.Message() != "" {
		t.Fatal("Invalid error message. ")
	}

	if len(err.Fields()) != 0 {
		t.Fatal("Invalid error fields. ")
	}

	if field := err.Fields().Get("test"); field.Key != "" || field.Message != "" {
		t.Fatal("Invalid error fields. ")
	}
}
