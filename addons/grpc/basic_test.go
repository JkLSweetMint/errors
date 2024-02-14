package sm_errors_grpc

import (
	"errors"
	"errors/entities"
	"testing"
)

func TestBasicError(t *testing.T) {
	var err = ExampleBasicGrpcError.SetError(errors.New("Test basic-grpc error. "))

	if v := err.ID(); 0000000000000001 != v {
		t.Fatalf("Invalid error id: '%d'. ", v)
	}

	if v := err.Status(); entities.StatusFailed != v {
		t.Fatalf("Invalid error status: '%s'. ", v)
	}

	if v := err.Message(); "Example basic-grpc message. " != v {
		t.Fatalf("Invalid error message: '%s'. ", v)
	}

	if v := err.StatusCode(); v != StatusInternalServerError {
		t.Fatalf("Invalid error grpc status code: '%s'. ", v)
	}

	if v := err.StatusCode().String(); v != StatusInternalServerError.String() {
		t.Fatalf("Invalid error grpc status code: '%s'. ", v)
	}
}
