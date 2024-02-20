package sm_errors_http

import (
	"errors"
	"sm_errors/entities"
	"testing"
)

func TestFieldsError(t *testing.T) {
	var err = ExampleFieldsHttpError.SetError(errors.New("Test fields-http error. "))

	if v := err.ID(); 0000000000000001 != v {
		t.Fatalf("Invalid error id: '%d'. ", v)
	}

	if v := err.Status(); entities.StatusFailed != v {
		t.Fatalf("Invalid error status: '%s'. ", v)
	}

	if v := err.Message(); "Example fields-http message. " != v {
		t.Fatalf("Invalid error message: '%s'. ", v)
	}

	if v := err.StatusCode(); v != StatusInternalServerError {
		t.Fatalf("Invalid error http status code: '%s'. ", v)
	}

	if v := err.StatusCode().String(); v != StatusInternalServerError.String() {
		t.Fatalf("Invalid error http status code: '%s'. ", v)
	}

	if fields := err.Fields(); fields == nil || len(fields) != 1 {
		t.Fatalf("Invalid error fields: '%s'. ", fields)
	} else {
		var field = fields[0]

		if v := field.Key; v != "login" {
			t.Fatalf("Invalid error field key: '%s'. ", v)
		}

		if v := field.Message; v != "Is required. " {
			t.Fatalf("Invalid error field message: '%s'. ", v)
		}
	}
}
