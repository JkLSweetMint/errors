package sm_errors_http

import (
	"errors"
	"errors/entities"
	"testing"
)

func TestBasicError(t *testing.T) {
	var err = ExampleBasicHttpError.SetError(errors.New("Test basic-http error. "))

	if v := err.ID(); 0000000000000001 != v {
		t.Fatalf("Invalid error id: '%d'. ", v)
	}

	if v := err.Status(); entities.StatusFailed != v {
		t.Fatalf("Invalid error status: '%s'. ", v)
	}

	if v := err.Message(); "Example basic-http message. " != v {
		t.Fatalf("Invalid error message: '%s'. ", v)
	}

	if v := err.StatusCode(); v != StatusInternalServerError {
		t.Fatalf("Invalid error http status code: '%s'. ", v)
	}

	if v := err.StatusCode().String(); v != StatusInternalServerError.String() {
		t.Fatalf("Invalid error http status code: '%s'. ", v)
	}
}
