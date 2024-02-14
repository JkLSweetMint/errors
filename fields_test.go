package sm_errors

import (
	"errors"
	"errors/entities"
	"testing"
)

func TestFieldsError(t *testing.T) {
	var err = ExampleFieldsError.SetError(errors.New("Test fields error. "))

	if v := err.ID(); 0000000000000001 != v {
		t.Fatalf("Invalid error id: '%d'. ", v)
	}

	if v := err.Status(); entities.StatusFailed != v {
		t.Fatalf("Invalid error status: '%s'. ", v)
	}

	if v := err.Message(); "Example fields message. " != v {
		t.Fatalf("Invalid error message: '%s'. ", v)
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

func TestFieldsErrorWithMessageFormatting(t *testing.T) {
	var err = ExampleFieldsErrorWithMessageFormat.SetError(errors.New("Test basic error. "))

	if v := err.ID(); 0000000000000001 != v {
		t.Fatalf("Invalid error id: '%d'. ", v)
	}

	if v := err.Status(); entities.StatusFailed != v {
		t.Fatalf("Invalid error status: '%s'. ", v)
	}

	if v := err.Message(entities.MessageOption{
		Key:   "text",
		Value: "123",
	}); "Example fields message with message formatting text=123. " != v {
		t.Fatalf("Invalid error message: '%s'. ", v)
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

func TestFieldsErrorWithoutMessage(t *testing.T) {
	var err = ExampleFieldsError.SetError(errors.New("Test fields error. "))

	err.(*fields).message = nil

	if v := err.ID(); 0000000000000001 != v {
		t.Fatalf("Invalid error id: '%d'. ", v)
	}

	if v := err.Status(); entities.StatusFailed != v {
		t.Fatalf("Invalid error status: '%s'. ", v)
	}

	if v := err.Message(); "Test fields error. " != v {
		t.Fatalf("Invalid error message: '%s'. ", v)
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
