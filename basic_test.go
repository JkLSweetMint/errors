package sm_errors

import (
	"errors"
	"sm_errors/entities"
	"testing"
)

func TestBasicError(t *testing.T) {
	var err = ExampleBasicError.SetError(errors.New("Test basic error. "))

	if v := err.ID(); 0000000000000001 != v {
		t.Fatalf("Invalid error id: '%d'. ", v)
	}

	if v := err.Status(); entities.StatusFailed != v {
		t.Fatalf("Invalid error status: '%s'. ", v)
	}

	if v := err.Message(); "Example basic message. " != v {
		t.Fatalf("Invalid error message: '%s'. ", v)
	}
}

func TestBasicErrorWithMessageFormatting(t *testing.T) {
	var err = ExampleBasicErrorWithMessageFormat.SetError(errors.New("Test basic error. "))

	if v := err.ID(); 0000000000000001 != v {
		t.Fatalf("Invalid error id: '%d'. ", v)
	}

	if v := err.Status(); entities.StatusFailed != v {
		t.Fatalf("Invalid error status: '%s'. ", v)
	}

	if v := err.Message(entities.MessageOption{
		Key:   "text",
		Value: "123",
	}); "Example basic message with message formatting text=123. " != v {
		t.Fatalf("Invalid error message: '%s'. ", v)
	}
}

func TestBasicErrorWithoutMessage(t *testing.T) {
	var err = ExampleBasicError.SetError(errors.New("Test basic error. "))

	err.(*basic).message = nil

	if v := err.ID(); 0000000000000001 != v {
		t.Fatalf("Invalid error id: '%d'. ", v)
	}

	if v := err.Status(); entities.StatusFailed != v {
		t.Fatalf("Invalid error status: '%s'. ", v)
	}

	if v := err.Message(); "Test basic error. " != v {
		t.Fatalf("Invalid error message: '%s'. ", v)
	}
}
