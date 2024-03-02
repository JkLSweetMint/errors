package errors

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"gopkg.in/yaml.v3"
	"slices"
	grpc_addon "sm_errors/addons/grpc"
	web_addon "sm_errors/addons/web"
	web_http_addon "sm_errors/addons/web/http"
	web_ws_addon "sm_errors/addons/web/ws"
	"sm_errors/entities"
	"testing"
)

var (
	ExampleError = Constructor{
		ID:      1000000000000000,
		Status:  entities.StatusFailed,
		Message: new(entities.Message).Set("Example error message. "),

		Grpc: &grpc_addon.Constructor{
			StatusCode: grpc_addon.StatusInternalServerError,
		},
		Web: &web_addon.Constructor{
			Http: &web_http_addon.Constructor{
				StatusCode: web_http_addon.StatusInternalServerError,
			},
			Ws: &web_ws_addon.Constructor{
				StatusCode: web_ws_addon.StatusInternalServerError,
			},
		},
	}

	ExampleErrorEmptyMessage = Constructor{
		ID:     1000000000000001,
		Status: entities.StatusFailed,

		Grpc: &grpc_addon.Constructor{
			StatusCode: grpc_addon.StatusInternalServerError,
		},
		Web: &web_addon.Constructor{
			Http: &web_http_addon.Constructor{
				StatusCode: web_http_addon.StatusInternalServerError,
			},
			Ws: &web_ws_addon.Constructor{
				StatusCode: web_ws_addon.StatusInternalServerError,
			},
		},
	}

	ExampleFieldsError = Constructor{
		ID:      1000000000000000,
		Status:  entities.StatusFailed,
		Message: new(entities.Message).Set("Example fields error message. "),

		Grpc: &grpc_addon.Constructor{
			StatusCode: grpc_addon.StatusInternalServerError,
		},
		Web: &web_addon.Constructor{
			Http: &web_http_addon.Constructor{
				StatusCode: web_http_addon.StatusInternalServerError,
			},
			Ws: &web_ws_addon.Constructor{
				StatusCode: web_ws_addon.StatusInternalServerError,
			},
		},
	}

	ExampleFieldsErrorEmptyMessage = Constructor{
		ID:     1000000000000001,
		Status: entities.StatusFailed,

		Grpc: &grpc_addon.Constructor{
			StatusCode: grpc_addon.StatusInternalServerError,
		},
		Web: &web_addon.Constructor{
			Http: &web_http_addon.Constructor{
				StatusCode: web_http_addon.StatusInternalServerError,
			},
			Ws: &web_ws_addon.Constructor{
				StatusCode: web_ws_addon.StatusInternalServerError,
			},
		},
	}
)

// ----------------------- Logic -----------------------

func TestBasicError(t *testing.T) {
	// Step 1
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleError.SetError(srcErr).Build().ToBasic()
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
	}

	// Step 2
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleErrorEmptyMessage.SetError(srcErr).Build().ToBasic()
		)

		if err.ID() != 1000000000000001 {
			t.Fatal("Invalid error id. ")
		}

		if err.String() != "Test error. " {
			t.Fatal("Invalid string error. ")
		}

		if err.Message() != "Test error. " {
			t.Fatal("Invalid error message. ")
		}
	}
}

func TestGrpcError(t *testing.T) {
	// Step 1
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleError.SetError(srcErr).Build().ToGrpc()
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

		if err.StatusCode() != grpc_addon.StatusInternalServerError {
			t.Fatal("Invalid error grpc status code. ")
		}
	}

	// Step 2
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleErrorEmptyMessage.SetError(srcErr).Build().ToGrpc()
		)

		if err.ID() != 1000000000000001 {
			t.Fatal("Invalid error id. ")
		}

		if err.String() != "Test error. " {
			t.Fatal("Invalid string error. ")
		}

		if err.Message() != "Test error. " {
			t.Fatal("Invalid error message. ")
		}
	}
}

func TestWebError(t *testing.T) {
	// Step 1
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleError.SetError(srcErr).Build().ToWeb()
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
	}

	// Step 2
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleErrorEmptyMessage.SetError(srcErr).Build().ToWeb()
		)

		if err.ID() != 1000000000000001 {
			t.Fatal("Invalid error id. ")
		}

		if err.String() != "Test error. " {
			t.Fatal("Invalid string error. ")
		}

		if err.Message() != "Test error. " {
			t.Fatal("Invalid error message. ")
		}
	}
}

func TestWebHttpError(t *testing.T) {
	// Step 1
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleError.SetError(srcErr).Build().ToWeb().ToHttp()
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

		if err.StatusCode() != web_http_addon.StatusInternalServerError {
			t.Fatal("Invalid error http status code. ")
		}
	}

	// Step 2
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleErrorEmptyMessage.SetError(srcErr).Build().ToWeb().ToHttp()
		)

		if err.ID() != 1000000000000001 {
			t.Fatal("Invalid error id. ")
		}

		if err.String() != "Test error. " {
			t.Fatal("Invalid string error. ")
		}

		if err.Message() != "Test error. " {
			t.Fatal("Invalid error message. ")
		}
	}
}

func TestWebWsError(t *testing.T) {
	// Step 1
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleError.SetError(srcErr).Build().ToWeb().ToWs()
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

		if err.StatusCode() != web_ws_addon.StatusInternalServerError {
			t.Fatal("Invalid error ws status code. ")
		}
	}

	// Step 2
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleErrorEmptyMessage.SetError(srcErr).Build().ToWeb().ToWs()
		)

		if err.ID() != 1000000000000001 {
			t.Fatal("Invalid error id. ")
		}

		if err.String() != "Test error. " {
			t.Fatal("Invalid string error. ")
		}

		if err.Message() != "Test error. " {
			t.Fatal("Invalid error message. ")
		}
	}
}

func TestBasicFieldsError(t *testing.T) {
	// Step 1
	{
		var (
			srcErr = errors.New("Test fields error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToBasic()
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

		if len(err.Fields()) != 1 {
			t.Fatal("Invalid error fields. ")
		}

		if field := err.Fields().Get("test"); field.Key != "test" || field.Message != "Invalid value" {
			t.Fatal("Invalid error fields. ")
		}
	}

	// Step 2
	{
		var (
			srcErr = errors.New("Test fields error. ")
			err    = ExampleErrorEmptyMessage.SetError(srcErr).Build().ToBasic()
		)

		if err.ID() != 1000000000000001 {
			t.Fatal("Invalid error id. ")
		}

		if err.String() != "Test fields error. " {
			t.Fatal("Invalid string error. ")
		}

		if err.Message() != "Test fields error. " {
			t.Fatal("Invalid error message. ")
		}

		if len(err.Fields()) != 0 {
			t.Fatal("Invalid error fields. ")
		}

		if field := err.Fields().Get("test"); field.Key != "" || field.Message != "" {
			t.Fatal("Invalid error fields. ")
		}
	}
}

func TestGrpcFieldsError(t *testing.T) {
	// Step 1
	{
		var (
			srcErr = errors.New("Test fields error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToGrpc()
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

		if err.StatusCode() != grpc_addon.StatusInternalServerError {
			t.Fatal("Invalid error grpc status code. ")
		}

		if len(err.Fields()) != 1 {
			t.Fatal("Invalid error fields. ")
		}

		if field := err.Fields().Get("test"); field.Key != "test" || field.Message != "Invalid value" {
			t.Fatal("Invalid error fields. ")
		}
	}

	// Step 2
	{
		var (
			srcErr = errors.New("Test fields error. ")
			err    = ExampleFieldsErrorEmptyMessage.SetError(srcErr).Build().ToGrpc()
		)

		if err.ID() != 1000000000000001 {
			t.Fatal("Invalid error id. ")
		}

		if err.String() != "Test fields error. " {
			t.Fatal("Invalid string error. ")
		}

		if err.Message() != "Test fields error. " {
			t.Fatal("Invalid error message. ")
		}

		if len(err.Fields()) != 0 {
			t.Fatal("Invalid error fields. ")
		}

		if field := err.Fields().Get("test"); field.Key != "" || field.Message != "" {
			t.Fatal("Invalid error fields. ")
		}
	}
}

func TestWebFieldsError(t *testing.T) {
	// Step 1
	{
		var (
			srcErr = errors.New("Test fields error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToWeb().ToWs()
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

		if err.StatusCode() != web_ws_addon.StatusInternalServerError {
			t.Fatal("Invalid error http status code. ")
		}

		if len(err.Fields()) != 1 {
			t.Fatal("Invalid error fields. ")
		}

		if field := err.Fields().Get("test"); field.Key != "test" || field.Message != "Invalid value" {
			t.Fatal("Invalid error fields. ")
		}
	}

	// Step 2
	{
		var (
			srcErr = errors.New("Test fields error. ")
			err    = ExampleFieldsErrorEmptyMessage.SetError(srcErr).Build().ToWeb().ToHttp()
		)

		if err.ID() != 1000000000000001 {
			t.Fatal("Invalid error id. ")
		}

		if err.String() != "Test fields error. " {
			t.Fatal("Invalid string error. ")
		}

		if err.Message() != "Test fields error. " {
			t.Fatal("Invalid error message. ")
		}

		if len(err.Fields()) != 0 {
			t.Fatal("Invalid error fields. ")
		}

		if field := err.Fields().Get("test"); field.Key != "" || field.Message != "" {
			t.Fatal("Invalid error fields. ")
		}
	}
}

func TestWebHttpFieldsError(t *testing.T) {
	// Step 1
	{
		var (
			srcErr = errors.New("Test fields error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToWeb().ToWs()
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

		if err.StatusCode() != web_ws_addon.StatusInternalServerError {
			t.Fatal("Invalid error http status code. ")
		}

		if len(err.Fields()) != 1 {
			t.Fatal("Invalid error fields. ")
		}

		if field := err.Fields().Get("test"); field.Key != "test" || field.Message != "Invalid value" {
			t.Fatal("Invalid error fields. ")
		}
	}

	// Step 2
	{
		var (
			srcErr = errors.New("Test fields error. ")
			err    = ExampleFieldsErrorEmptyMessage.SetError(srcErr).Build().ToWeb().ToHttp()
		)

		if err.ID() != 1000000000000001 {
			t.Fatal("Invalid error id. ")
		}

		if err.String() != "Test fields error. " {
			t.Fatal("Invalid string error. ")
		}

		if err.Message() != "Test fields error. " {
			t.Fatal("Invalid error message. ")
		}

		if len(err.Fields()) != 0 {
			t.Fatal("Invalid error fields. ")
		}

		if field := err.Fields().Get("test"); field.Key != "" || field.Message != "" {
			t.Fatal("Invalid error fields. ")
		}
	}
}

func TestWebWsFieldsError(t *testing.T) {
	// Step 1
	{
		var (
			srcErr = errors.New("Test fields error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToWeb().ToWs()
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

		if err.StatusCode() != web_ws_addon.StatusInternalServerError {
			t.Fatal("Invalid error ws status code. ")
		}

		if len(err.Fields()) != 1 {
			t.Fatal("Invalid error fields. ")
		}

		if field := err.Fields().Get("test"); field.Key != "test" || field.Message != "Invalid value" {
			t.Fatal("Invalid error fields. ")
		}
	}

	// Step 2
	{
		var (
			srcErr = errors.New("Test fields error. ")
			err    = ExampleFieldsErrorEmptyMessage.SetError(srcErr).Build().ToWeb().ToWs()
		)

		if err.ID() != 1000000000000001 {
			t.Fatal("Invalid error id. ")
		}

		if err.String() != "Test fields error. " {
			t.Fatal("Invalid string error. ")
		}

		if err.Message() != "Test fields error. " {
			t.Fatal("Invalid error message. ")
		}

		if len(err.Fields()) != 0 {
			t.Fatal("Invalid error fields. ")
		}

		if field := err.Fields().Get("test"); field.Key != "" || field.Message != "" {
			t.Fatal("Invalid error fields. ")
		}
	}
}

// ----------------------- Marshaling -----------------------

func TestMarshalBasicError(t *testing.T) {
	// JSON
	{
		var (
			srcErr         = errors.New("Test error. ")
			err            = ExampleError.SetError(srcErr).Build().ToBasic()
			exampleMarshal = []byte{123, 34, 101, 114, 114, 111, 114, 34, 58, 123, 34, 105, 100, 34, 58, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 98, 97, 115, 105, 99, 34, 125, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 69, 120, 97, 109, 112, 108, 101, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 105, 108, 101, 100, 34, 125}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = json.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in json: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in json. ")
		}
	}

	// YAML
	{
		var (
			srcErr         = errors.New("Test error. ")
			err            = ExampleError.SetError(srcErr).Build().ToBasic()
			exampleMarshal = []byte{101, 114, 114, 111, 114, 58, 10, 32, 32, 32, 32, 105, 100, 58, 32, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 10, 32, 32, 32, 32, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 39, 10, 32, 32, 32, 32, 116, 121, 112, 101, 58, 32, 98, 97, 115, 105, 99, 10, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 69, 120, 97, 109, 112, 108, 101, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 39, 10, 115, 116, 97, 116, 117, 115, 58, 32, 102, 97, 105, 108, 101, 100, 10}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = yaml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in yaml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in yaml. ")
		}
	}

	// XML
	{
		var (
			srcErr         = errors.New("Test error. ")
			err            = ExampleError.SetError(srcErr).Build().ToBasic()
			exampleMarshal = []byte{60, 101, 114, 114, 111, 114, 62, 60, 105, 100, 62, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 60, 47, 105, 100, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 116, 121, 112, 101, 62, 98, 97, 115, 105, 99, 60, 47, 116, 121, 112, 101, 62, 60, 47, 101, 114, 114, 111, 114, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 69, 120, 97, 109, 112, 108, 101, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 115, 116, 97, 116, 117, 115, 62, 102, 97, 105, 108, 101, 100, 60, 47, 115, 116, 97, 116, 117, 115, 62}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = xml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in xml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in xml. ")
		}
	}
}

func TestMarshalGrpcError(t *testing.T) {
	// JSON
	{
		var (
			srcErr         = errors.New("Test error. ")
			err            = ExampleError.SetError(srcErr).Build().ToGrpc()
			exampleMarshal = []byte{123, 34, 99, 111, 100, 101, 34, 58, 49, 51, 44, 34, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 73, 110, 116, 101, 114, 110, 97, 108, 32, 115, 101, 114, 118, 101, 114, 32, 101, 114, 114, 111, 114, 34, 44, 34, 101, 114, 114, 111, 114, 34, 58, 123, 34, 105, 100, 34, 58, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 98, 97, 115, 105, 99, 34, 125, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 69, 120, 97, 109, 112, 108, 101, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 105, 108, 101, 100, 34, 125}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = json.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in json: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in json. ")
		}
	}

	// YAML
	{
		var (
			srcErr         = errors.New("Test error. ")
			err            = ExampleError.SetError(srcErr).Build().ToGrpc()
			exampleMarshal = []byte{99, 111, 100, 101, 58, 32, 49, 51, 10, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 58, 32, 73, 110, 116, 101, 114, 110, 97, 108, 32, 115, 101, 114, 118, 101, 114, 32, 101, 114, 114, 111, 114, 10, 101, 114, 114, 111, 114, 58, 10, 32, 32, 32, 32, 105, 100, 58, 32, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 10, 32, 32, 32, 32, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 39, 10, 32, 32, 32, 32, 116, 121, 112, 101, 58, 32, 98, 97, 115, 105, 99, 10, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 69, 120, 97, 109, 112, 108, 101, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 39, 10, 115, 116, 97, 116, 117, 115, 58, 32, 102, 97, 105, 108, 101, 100, 10}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = yaml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in yaml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in yaml. ")
		}
	}

	// XML
	{
		var (
			srcErr         = errors.New("Test error. ")
			err            = ExampleError.SetError(srcErr).Build().ToGrpc()
			exampleMarshal = []byte{60, 99, 111, 100, 101, 62, 73, 110, 116, 101, 114, 110, 97, 108, 32, 115, 101, 114, 118, 101, 114, 32, 101, 114, 114, 111, 114, 60, 47, 99, 111, 100, 101, 62, 60, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 62, 73, 110, 116, 101, 114, 110, 97, 108, 32, 115, 101, 114, 118, 101, 114, 32, 101, 114, 114, 111, 114, 60, 47, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 62, 60, 101, 114, 114, 111, 114, 62, 60, 105, 100, 62, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 60, 47, 105, 100, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 116, 121, 112, 101, 62, 98, 97, 115, 105, 99, 60, 47, 116, 121, 112, 101, 62, 60, 47, 101, 114, 114, 111, 114, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 69, 120, 97, 109, 112, 108, 101, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 115, 116, 97, 116, 117, 115, 62, 102, 97, 105, 108, 101, 100, 60, 47, 115, 116, 97, 116, 117, 115, 62}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = xml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in xml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in xml. ")
		}
	}
}

func TestMarshalWebError(t *testing.T) {
	// JSON
	{
		var (
			srcErr         = errors.New("Test error. ")
			err            = ExampleError.SetError(srcErr).Build().ToWeb()
			exampleMarshal = []byte{123, 34, 101, 114, 114, 111, 114, 34, 58, 123, 34, 105, 100, 34, 58, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 98, 97, 115, 105, 99, 34, 125, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 69, 120, 97, 109, 112, 108, 101, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 105, 108, 101, 100, 34, 125}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = json.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in json: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in json. ")
		}
	}

	// YAML
	{
		var (
			srcErr         = errors.New("Test error. ")
			err            = ExampleError.SetError(srcErr).Build().ToWeb()
			exampleMarshal = []byte{101, 114, 114, 111, 114, 58, 10, 32, 32, 32, 32, 105, 100, 58, 32, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 10, 32, 32, 32, 32, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 39, 10, 32, 32, 32, 32, 116, 121, 112, 101, 58, 32, 98, 97, 115, 105, 99, 10, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 69, 120, 97, 109, 112, 108, 101, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 39, 10, 115, 116, 97, 116, 117, 115, 58, 32, 102, 97, 105, 108, 101, 100, 10}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = yaml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in yaml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in yaml. ")
		}
	}

	// XML
	{
		var (
			srcErr         = errors.New("Test error. ")
			err            = ExampleError.SetError(srcErr).Build().ToWeb()
			exampleMarshal = []byte{60, 101, 114, 114, 111, 114, 62, 60, 105, 100, 62, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 60, 47, 105, 100, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 116, 121, 112, 101, 62, 98, 97, 115, 105, 99, 60, 47, 116, 121, 112, 101, 62, 60, 47, 101, 114, 114, 111, 114, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 69, 120, 97, 109, 112, 108, 101, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 115, 116, 97, 116, 117, 115, 62, 102, 97, 105, 108, 101, 100, 60, 47, 115, 116, 97, 116, 117, 115, 62}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = xml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in xml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in xml. ")
		}
	}
}

func TestMarshalWebHttpError(t *testing.T) {
	// JSON
	{
		var (
			srcErr         = errors.New("Test error. ")
			err            = ExampleError.SetError(srcErr).Build().ToWeb().ToHttp()
			exampleMarshal = []byte{123, 34, 99, 111, 100, 101, 34, 58, 53, 48, 48, 44, 34, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 73, 110, 116, 101, 114, 110, 97, 108, 32, 83, 101, 114, 118, 101, 114, 32, 69, 114, 114, 111, 114, 34, 44, 34, 101, 114, 114, 111, 114, 34, 58, 123, 34, 105, 100, 34, 58, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 98, 97, 115, 105, 99, 34, 125, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 69, 120, 97, 109, 112, 108, 101, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 105, 108, 101, 100, 34, 125}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = json.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in json: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in json. ")
		}
	}

	// YAML
	{
		var (
			srcErr         = errors.New("Test error. ")
			err            = ExampleError.SetError(srcErr).Build().ToWeb().ToHttp()
			exampleMarshal = []byte{99, 111, 100, 101, 58, 32, 53, 48, 48, 10, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 58, 32, 73, 110, 116, 101, 114, 110, 97, 108, 32, 83, 101, 114, 118, 101, 114, 32, 69, 114, 114, 111, 114, 10, 101, 114, 114, 111, 114, 58, 10, 32, 32, 32, 32, 105, 100, 58, 32, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 10, 32, 32, 32, 32, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 39, 10, 32, 32, 32, 32, 116, 121, 112, 101, 58, 32, 98, 97, 115, 105, 99, 10, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 69, 120, 97, 109, 112, 108, 101, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 39, 10, 115, 116, 97, 116, 117, 115, 58, 32, 102, 97, 105, 108, 101, 100, 10}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = yaml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in yaml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in yaml. ")
		}
	}

	// XML
	{
		var (
			srcErr         = errors.New("Test error. ")
			err            = ExampleError.SetError(srcErr).Build().ToWeb().ToHttp()
			exampleMarshal = []byte{60, 99, 111, 100, 101, 62, 73, 110, 116, 101, 114, 110, 97, 108, 32, 83, 101, 114, 118, 101, 114, 32, 69, 114, 114, 111, 114, 60, 47, 99, 111, 100, 101, 62, 60, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 62, 73, 110, 116, 101, 114, 110, 97, 108, 32, 83, 101, 114, 118, 101, 114, 32, 69, 114, 114, 111, 114, 60, 47, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 62, 60, 101, 114, 114, 111, 114, 62, 60, 105, 100, 62, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 60, 47, 105, 100, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 116, 121, 112, 101, 62, 98, 97, 115, 105, 99, 60, 47, 116, 121, 112, 101, 62, 60, 47, 101, 114, 114, 111, 114, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 69, 120, 97, 109, 112, 108, 101, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 115, 116, 97, 116, 117, 115, 62, 102, 97, 105, 108, 101, 100, 60, 47, 115, 116, 97, 116, 117, 115, 62}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = xml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in xml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in xml. ")
		}
	}
}

func TestMarshalWebWsError(t *testing.T) {
	// JSON
	{
		var (
			srcErr         = errors.New("Test error. ")
			err            = ExampleError.SetError(srcErr).Build().ToWeb().ToWs()
			exampleMarshal = []byte{123, 34, 99, 111, 100, 101, 34, 58, 49, 48, 49, 49, 44, 34, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 73, 110, 116, 101, 114, 110, 97, 108, 32, 115, 101, 114, 118, 101, 114, 32, 101, 114, 114, 111, 114, 32, 119, 104, 105, 108, 101, 32, 111, 112, 101, 114, 97, 116, 105, 110, 103, 34, 44, 34, 101, 114, 114, 111, 114, 34, 58, 123, 34, 105, 100, 34, 58, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 98, 97, 115, 105, 99, 34, 125, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 69, 120, 97, 109, 112, 108, 101, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 105, 108, 101, 100, 34, 125}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = json.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in json: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in json. ")
		}
	}

	// YAML
	{
		var (
			srcErr         = errors.New("Test error. ")
			err            = ExampleError.SetError(srcErr).Build().ToWeb().ToWs()
			exampleMarshal = []byte{99, 111, 100, 101, 58, 32, 49, 48, 49, 49, 10, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 58, 32, 73, 110, 116, 101, 114, 110, 97, 108, 32, 115, 101, 114, 118, 101, 114, 32, 101, 114, 114, 111, 114, 32, 119, 104, 105, 108, 101, 32, 111, 112, 101, 114, 97, 116, 105, 110, 103, 10, 101, 114, 114, 111, 114, 58, 10, 32, 32, 32, 32, 105, 100, 58, 32, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 10, 32, 32, 32, 32, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 39, 10, 32, 32, 32, 32, 116, 121, 112, 101, 58, 32, 98, 97, 115, 105, 99, 10, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 69, 120, 97, 109, 112, 108, 101, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 39, 10, 115, 116, 97, 116, 117, 115, 58, 32, 102, 97, 105, 108, 101, 100, 10}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = yaml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in yaml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in yaml. ")
		}
	}

	// XML
	{
		var (
			srcErr         = errors.New("Test error. ")
			err            = ExampleError.SetError(srcErr).Build().ToWeb().ToWs()
			exampleMarshal = []byte{60, 99, 111, 100, 101, 62, 73, 110, 116, 101, 114, 110, 97, 108, 32, 115, 101, 114, 118, 101, 114, 32, 101, 114, 114, 111, 114, 32, 119, 104, 105, 108, 101, 32, 111, 112, 101, 114, 97, 116, 105, 110, 103, 60, 47, 99, 111, 100, 101, 62, 60, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 62, 73, 110, 116, 101, 114, 110, 97, 108, 32, 115, 101, 114, 118, 101, 114, 32, 101, 114, 114, 111, 114, 32, 119, 104, 105, 108, 101, 32, 111, 112, 101, 114, 97, 116, 105, 110, 103, 60, 47, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 62, 60, 101, 114, 114, 111, 114, 62, 60, 105, 100, 62, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 60, 47, 105, 100, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 116, 121, 112, 101, 62, 98, 97, 115, 105, 99, 60, 47, 116, 121, 112, 101, 62, 60, 47, 101, 114, 114, 111, 114, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 69, 120, 97, 109, 112, 108, 101, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 115, 116, 97, 116, 117, 115, 62, 102, 97, 105, 108, 101, 100, 60, 47, 115, 116, 97, 116, 117, 115, 62}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = xml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in xml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in xml. ")
		}
	}
}

func TestMarshalBasicFieldsError(t *testing.T) {
	// JSON
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToBasic()
			exampleMarshal = []byte{123, 34, 101, 114, 114, 111, 114, 34, 58, 123, 34, 102, 105, 101, 108, 100, 115, 34, 58, 91, 123, 34, 75, 101, 121, 34, 58, 34, 116, 101, 115, 116, 34, 44, 34, 77, 101, 115, 115, 97, 103, 101, 34, 58, 34, 73, 110, 118, 97, 108, 105, 100, 32, 118, 97, 108, 117, 101, 34, 125, 93, 44, 34, 105, 100, 34, 58, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 102, 105, 101, 108, 100, 115, 34, 125, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 69, 120, 97, 109, 112, 108, 101, 32, 102, 105, 101, 108, 100, 115, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 105, 108, 101, 100, 34, 125}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = json.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in json: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in json. ")
		}
	}

	// YAML
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToBasic()
			exampleMarshal = []byte{101, 114, 114, 111, 114, 58, 10, 32, 32, 32, 32, 102, 105, 101, 108, 100, 115, 58, 10, 32, 32, 32, 32, 32, 32, 32, 32, 45, 32, 107, 101, 121, 58, 32, 116, 101, 115, 116, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 109, 101, 115, 115, 97, 103, 101, 58, 32, 73, 110, 118, 97, 108, 105, 100, 32, 118, 97, 108, 117, 101, 10, 32, 32, 32, 32, 105, 100, 58, 32, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 10, 32, 32, 32, 32, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 39, 10, 32, 32, 32, 32, 116, 121, 112, 101, 58, 32, 102, 105, 101, 108, 100, 115, 10, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 69, 120, 97, 109, 112, 108, 101, 32, 102, 105, 101, 108, 100, 115, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 39, 10, 115, 116, 97, 116, 117, 115, 58, 32, 102, 97, 105, 108, 101, 100, 10}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = yaml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in yaml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in yaml. ")
		}
	}

	// XML
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToBasic()
			exampleMarshal = []byte{60, 101, 114, 114, 111, 114, 62, 60, 102, 105, 101, 108, 100, 115, 62, 91, 123, 116, 101, 115, 116, 32, 73, 110, 118, 97, 108, 105, 100, 32, 118, 97, 108, 117, 101, 125, 93, 60, 47, 102, 105, 101, 108, 100, 115, 62, 60, 105, 100, 62, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 60, 47, 105, 100, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 116, 121, 112, 101, 62, 102, 105, 101, 108, 100, 115, 60, 47, 116, 121, 112, 101, 62, 60, 47, 101, 114, 114, 111, 114, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 69, 120, 97, 109, 112, 108, 101, 32, 102, 105, 101, 108, 100, 115, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 115, 116, 97, 116, 117, 115, 62, 102, 97, 105, 108, 101, 100, 60, 47, 115, 116, 97, 116, 117, 115, 62}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = xml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in xml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in xml. ")
		}
	}
}

func TestMarshalGrpcFieldsError(t *testing.T) {
	// JSON
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToGrpc()
			exampleMarshal = []byte{123, 34, 99, 111, 100, 101, 34, 58, 49, 51, 44, 34, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 73, 110, 116, 101, 114, 110, 97, 108, 32, 115, 101, 114, 118, 101, 114, 32, 101, 114, 114, 111, 114, 34, 44, 34, 101, 114, 114, 111, 114, 34, 58, 123, 34, 102, 105, 101, 108, 100, 115, 34, 58, 91, 123, 34, 75, 101, 121, 34, 58, 34, 116, 101, 115, 116, 34, 44, 34, 77, 101, 115, 115, 97, 103, 101, 34, 58, 34, 73, 110, 118, 97, 108, 105, 100, 32, 118, 97, 108, 117, 101, 34, 125, 93, 44, 34, 105, 100, 34, 58, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 102, 105, 101, 108, 100, 115, 34, 125, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 69, 120, 97, 109, 112, 108, 101, 32, 102, 105, 101, 108, 100, 115, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 105, 108, 101, 100, 34, 125}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = json.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in json: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in json. ")
		}
	}

	// YAML
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToGrpc()
			exampleMarshal = []byte{99, 111, 100, 101, 58, 32, 49, 51, 10, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 58, 32, 73, 110, 116, 101, 114, 110, 97, 108, 32, 115, 101, 114, 118, 101, 114, 32, 101, 114, 114, 111, 114, 10, 101, 114, 114, 111, 114, 58, 10, 32, 32, 32, 32, 102, 105, 101, 108, 100, 115, 58, 10, 32, 32, 32, 32, 32, 32, 32, 32, 45, 32, 107, 101, 121, 58, 32, 116, 101, 115, 116, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 109, 101, 115, 115, 97, 103, 101, 58, 32, 73, 110, 118, 97, 108, 105, 100, 32, 118, 97, 108, 117, 101, 10, 32, 32, 32, 32, 105, 100, 58, 32, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 10, 32, 32, 32, 32, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 39, 10, 32, 32, 32, 32, 116, 121, 112, 101, 58, 32, 102, 105, 101, 108, 100, 115, 10, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 69, 120, 97, 109, 112, 108, 101, 32, 102, 105, 101, 108, 100, 115, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 39, 10, 115, 116, 97, 116, 117, 115, 58, 32, 102, 97, 105, 108, 101, 100, 10}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = yaml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in yaml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in yaml. ")
		}
	}

	// XML
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToGrpc()
			exampleMarshal = []byte{60, 99, 111, 100, 101, 62, 73, 110, 116, 101, 114, 110, 97, 108, 32, 115, 101, 114, 118, 101, 114, 32, 101, 114, 114, 111, 114, 60, 47, 99, 111, 100, 101, 62, 60, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 62, 73, 110, 116, 101, 114, 110, 97, 108, 32, 115, 101, 114, 118, 101, 114, 32, 101, 114, 114, 111, 114, 60, 47, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 62, 60, 101, 114, 114, 111, 114, 62, 60, 102, 105, 101, 108, 100, 115, 62, 91, 123, 116, 101, 115, 116, 32, 73, 110, 118, 97, 108, 105, 100, 32, 118, 97, 108, 117, 101, 125, 93, 60, 47, 102, 105, 101, 108, 100, 115, 62, 60, 105, 100, 62, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 60, 47, 105, 100, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 116, 121, 112, 101, 62, 102, 105, 101, 108, 100, 115, 60, 47, 116, 121, 112, 101, 62, 60, 47, 101, 114, 114, 111, 114, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 69, 120, 97, 109, 112, 108, 101, 32, 102, 105, 101, 108, 100, 115, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 115, 116, 97, 116, 117, 115, 62, 102, 97, 105, 108, 101, 100, 60, 47, 115, 116, 97, 116, 117, 115, 62}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = xml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in xml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in xml. ")
		}
	}
}

func TestMarshalWebFieldsError(t *testing.T) {
	// JSON
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToWeb()
			exampleMarshal = []byte{123, 34, 101, 114, 114, 111, 114, 34, 58, 123, 34, 102, 105, 101, 108, 100, 115, 34, 58, 91, 123, 34, 75, 101, 121, 34, 58, 34, 116, 101, 115, 116, 34, 44, 34, 77, 101, 115, 115, 97, 103, 101, 34, 58, 34, 73, 110, 118, 97, 108, 105, 100, 32, 118, 97, 108, 117, 101, 34, 125, 93, 44, 34, 105, 100, 34, 58, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 102, 105, 101, 108, 100, 115, 34, 125, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 69, 120, 97, 109, 112, 108, 101, 32, 102, 105, 101, 108, 100, 115, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 105, 108, 101, 100, 34, 125}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = json.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in json: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in json. ")
		}
	}

	// YAML
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToWeb()
			exampleMarshal = []byte{101, 114, 114, 111, 114, 58, 10, 32, 32, 32, 32, 102, 105, 101, 108, 100, 115, 58, 10, 32, 32, 32, 32, 32, 32, 32, 32, 45, 32, 107, 101, 121, 58, 32, 116, 101, 115, 116, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 109, 101, 115, 115, 97, 103, 101, 58, 32, 73, 110, 118, 97, 108, 105, 100, 32, 118, 97, 108, 117, 101, 10, 32, 32, 32, 32, 105, 100, 58, 32, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 10, 32, 32, 32, 32, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 39, 10, 32, 32, 32, 32, 116, 121, 112, 101, 58, 32, 102, 105, 101, 108, 100, 115, 10, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 69, 120, 97, 109, 112, 108, 101, 32, 102, 105, 101, 108, 100, 115, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 39, 10, 115, 116, 97, 116, 117, 115, 58, 32, 102, 97, 105, 108, 101, 100, 10}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = yaml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in yaml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in yaml. ")
		}
	}

	// XML
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToWeb()
			exampleMarshal = []byte{60, 101, 114, 114, 111, 114, 62, 60, 102, 105, 101, 108, 100, 115, 62, 91, 123, 116, 101, 115, 116, 32, 73, 110, 118, 97, 108, 105, 100, 32, 118, 97, 108, 117, 101, 125, 93, 60, 47, 102, 105, 101, 108, 100, 115, 62, 60, 105, 100, 62, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 60, 47, 105, 100, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 116, 121, 112, 101, 62, 102, 105, 101, 108, 100, 115, 60, 47, 116, 121, 112, 101, 62, 60, 47, 101, 114, 114, 111, 114, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 69, 120, 97, 109, 112, 108, 101, 32, 102, 105, 101, 108, 100, 115, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 115, 116, 97, 116, 117, 115, 62, 102, 97, 105, 108, 101, 100, 60, 47, 115, 116, 97, 116, 117, 115, 62}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = xml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in xml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in xml. ")
		}
	}
}

func TestMarshalWebHttpFieldsError(t *testing.T) {
	// JSON
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToWeb().ToHttp()
			exampleMarshal = []byte{123, 34, 99, 111, 100, 101, 34, 58, 53, 48, 48, 44, 34, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 73, 110, 116, 101, 114, 110, 97, 108, 32, 83, 101, 114, 118, 101, 114, 32, 69, 114, 114, 111, 114, 34, 44, 34, 101, 114, 114, 111, 114, 34, 58, 123, 34, 102, 105, 101, 108, 100, 115, 34, 58, 91, 123, 34, 75, 101, 121, 34, 58, 34, 116, 101, 115, 116, 34, 44, 34, 77, 101, 115, 115, 97, 103, 101, 34, 58, 34, 73, 110, 118, 97, 108, 105, 100, 32, 118, 97, 108, 117, 101, 34, 125, 93, 44, 34, 105, 100, 34, 58, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 102, 105, 101, 108, 100, 115, 34, 125, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 69, 120, 97, 109, 112, 108, 101, 32, 102, 105, 101, 108, 100, 115, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 105, 108, 101, 100, 34, 125}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = json.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in json: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in json. ")
		}
	}

	// YAML
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToWeb().ToHttp()
			exampleMarshal = []byte{99, 111, 100, 101, 58, 32, 53, 48, 48, 10, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 58, 32, 73, 110, 116, 101, 114, 110, 97, 108, 32, 83, 101, 114, 118, 101, 114, 32, 69, 114, 114, 111, 114, 10, 101, 114, 114, 111, 114, 58, 10, 32, 32, 32, 32, 102, 105, 101, 108, 100, 115, 58, 10, 32, 32, 32, 32, 32, 32, 32, 32, 45, 32, 107, 101, 121, 58, 32, 116, 101, 115, 116, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 109, 101, 115, 115, 97, 103, 101, 58, 32, 73, 110, 118, 97, 108, 105, 100, 32, 118, 97, 108, 117, 101, 10, 32, 32, 32, 32, 105, 100, 58, 32, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 10, 32, 32, 32, 32, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 39, 10, 32, 32, 32, 32, 116, 121, 112, 101, 58, 32, 102, 105, 101, 108, 100, 115, 10, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 69, 120, 97, 109, 112, 108, 101, 32, 102, 105, 101, 108, 100, 115, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 39, 10, 115, 116, 97, 116, 117, 115, 58, 32, 102, 97, 105, 108, 101, 100, 10}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = yaml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in yaml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in yaml. ")
		}
	}

	// XML
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToWeb().ToHttp()
			exampleMarshal = []byte{60, 99, 111, 100, 101, 62, 73, 110, 116, 101, 114, 110, 97, 108, 32, 83, 101, 114, 118, 101, 114, 32, 69, 114, 114, 111, 114, 60, 47, 99, 111, 100, 101, 62, 60, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 62, 73, 110, 116, 101, 114, 110, 97, 108, 32, 83, 101, 114, 118, 101, 114, 32, 69, 114, 114, 111, 114, 60, 47, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 62, 60, 101, 114, 114, 111, 114, 62, 60, 102, 105, 101, 108, 100, 115, 62, 91, 123, 116, 101, 115, 116, 32, 73, 110, 118, 97, 108, 105, 100, 32, 118, 97, 108, 117, 101, 125, 93, 60, 47, 102, 105, 101, 108, 100, 115, 62, 60, 105, 100, 62, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 60, 47, 105, 100, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 116, 121, 112, 101, 62, 102, 105, 101, 108, 100, 115, 60, 47, 116, 121, 112, 101, 62, 60, 47, 101, 114, 114, 111, 114, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 69, 120, 97, 109, 112, 108, 101, 32, 102, 105, 101, 108, 100, 115, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 115, 116, 97, 116, 117, 115, 62, 102, 97, 105, 108, 101, 100, 60, 47, 115, 116, 97, 116, 117, 115, 62}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = xml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in xml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in xml. ")
		}
	}
}

func TestMarshalWebWsFieldsError(t *testing.T) {
	// JSON
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToWeb().ToWs()
			exampleMarshal = []byte{123, 34, 99, 111, 100, 101, 34, 58, 49, 48, 49, 49, 44, 34, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 73, 110, 116, 101, 114, 110, 97, 108, 32, 115, 101, 114, 118, 101, 114, 32, 101, 114, 114, 111, 114, 32, 119, 104, 105, 108, 101, 32, 111, 112, 101, 114, 97, 116, 105, 110, 103, 34, 44, 34, 101, 114, 114, 111, 114, 34, 58, 123, 34, 102, 105, 101, 108, 100, 115, 34, 58, 91, 123, 34, 75, 101, 121, 34, 58, 34, 116, 101, 115, 116, 34, 44, 34, 77, 101, 115, 115, 97, 103, 101, 34, 58, 34, 73, 110, 118, 97, 108, 105, 100, 32, 118, 97, 108, 117, 101, 34, 125, 93, 44, 34, 105, 100, 34, 58, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 34, 44, 34, 116, 121, 112, 101, 34, 58, 34, 102, 105, 101, 108, 100, 115, 34, 125, 44, 34, 109, 101, 115, 115, 97, 103, 101, 34, 58, 34, 69, 120, 97, 109, 112, 108, 101, 32, 102, 105, 101, 108, 100, 115, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 34, 44, 34, 115, 116, 97, 116, 117, 115, 34, 58, 34, 102, 97, 105, 108, 101, 100, 34, 125}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = json.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in json: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in json. ")
		}
	}

	// YAML
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToWeb().ToWs()
			exampleMarshal = []byte{99, 111, 100, 101, 58, 32, 49, 48, 49, 49, 10, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 58, 32, 73, 110, 116, 101, 114, 110, 97, 108, 32, 115, 101, 114, 118, 101, 114, 32, 101, 114, 114, 111, 114, 32, 119, 104, 105, 108, 101, 32, 111, 112, 101, 114, 97, 116, 105, 110, 103, 10, 101, 114, 114, 111, 114, 58, 10, 32, 32, 32, 32, 102, 105, 101, 108, 100, 115, 58, 10, 32, 32, 32, 32, 32, 32, 32, 32, 45, 32, 107, 101, 121, 58, 32, 116, 101, 115, 116, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 109, 101, 115, 115, 97, 103, 101, 58, 32, 73, 110, 118, 97, 108, 105, 100, 32, 118, 97, 108, 117, 101, 10, 32, 32, 32, 32, 105, 100, 58, 32, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 10, 32, 32, 32, 32, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 39, 10, 32, 32, 32, 32, 116, 121, 112, 101, 58, 32, 102, 105, 101, 108, 100, 115, 10, 109, 101, 115, 115, 97, 103, 101, 58, 32, 39, 69, 120, 97, 109, 112, 108, 101, 32, 102, 105, 101, 108, 100, 115, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 39, 10, 115, 116, 97, 116, 117, 115, 58, 32, 102, 97, 105, 108, 101, 100, 10}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = yaml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in yaml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in yaml. ")
		}
	}

	// XML
	{
		var (
			srcErr = errors.New("Test error. ")
			err    = ExampleFieldsError.SetError(srcErr).SetFields(entities.Field{
				Key:     "test",
				Message: "Invalid value",
			}).Build().ToWeb().ToWs()
			exampleMarshal = []byte{60, 99, 111, 100, 101, 62, 73, 110, 116, 101, 114, 110, 97, 108, 32, 115, 101, 114, 118, 101, 114, 32, 101, 114, 114, 111, 114, 32, 119, 104, 105, 108, 101, 32, 111, 112, 101, 114, 97, 116, 105, 110, 103, 60, 47, 99, 111, 100, 101, 62, 60, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 62, 73, 110, 116, 101, 114, 110, 97, 108, 32, 115, 101, 114, 118, 101, 114, 32, 101, 114, 114, 111, 114, 32, 119, 104, 105, 108, 101, 32, 111, 112, 101, 114, 97, 116, 105, 110, 103, 60, 47, 99, 111, 100, 101, 95, 109, 101, 115, 115, 97, 103, 101, 62, 60, 101, 114, 114, 111, 114, 62, 60, 102, 105, 101, 108, 100, 115, 62, 91, 123, 116, 101, 115, 116, 32, 73, 110, 118, 97, 108, 105, 100, 32, 118, 97, 108, 117, 101, 125, 93, 60, 47, 102, 105, 101, 108, 100, 115, 62, 60, 105, 100, 62, 49, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 60, 47, 105, 100, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 84, 101, 115, 116, 32, 101, 114, 114, 111, 114, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 116, 121, 112, 101, 62, 102, 105, 101, 108, 100, 115, 60, 47, 116, 121, 112, 101, 62, 60, 47, 101, 114, 114, 111, 114, 62, 60, 109, 101, 115, 115, 97, 103, 101, 62, 69, 120, 97, 109, 112, 108, 101, 32, 102, 105, 101, 108, 100, 115, 32, 101, 114, 114, 111, 114, 32, 109, 101, 115, 115, 97, 103, 101, 46, 32, 60, 47, 109, 101, 115, 115, 97, 103, 101, 62, 60, 115, 116, 97, 116, 117, 115, 62, 102, 97, 105, 108, 101, 100, 60, 47, 115, 116, 97, 116, 117, 115, 62}
		)

		var (
			data []byte
			err_ error
		)

		if data, err_ = xml.Marshal(err); err_ != nil {
			t.Fatalf("Failed to package the error in xml: '%v'. ", err_)
		}

		if !slices.Equal(data, exampleMarshal) {
			t.Fatal("Incorrect operation of error packaging in xml. ")
		}
	}
}
