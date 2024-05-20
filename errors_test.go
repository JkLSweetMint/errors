package errors

import (
	"encoding/xml"
	"errors"
	"fmt"
	"sm-errors/entities/details"
	"sm-errors/entities/messages"
	"sm-errors/types"
	"testing"
)

// Примеры базовой ошибки.
var (
	ExampleError = Constructor[Error]{
		ID:     "T-000001",
		Type:   types.TypeSystem,
		Status: types.StatusFatal,

		Message: new(messages.TextMessage).
			Text("Example error. "),
	}.Build()

	ExampleErrorWithDetails = Constructor[Error]{
		ID:     "T-000002",
		Type:   types.TypeSystem,
		Status: types.StatusFatal,

		Message: new(messages.TextMessage).
			Text("Example error with details. "),
		Details: new(details.Details).
			Set("key", "value"),
	}.Build()

	ExampleErrorWithDetailsAndFields = Constructor[Error]{
		ID:     "T-000003",
		Type:   types.TypeSystem,
		Status: types.StatusFatal,

		Message: new(messages.TextMessage).Text("Example error with details and fields. "),
		Details: new(details.Details).
			Set("key", "value").
			SetFields(types.DetailsField{
				Key:     new(details.FieldKey).Add("test"),
				Message: new(messages.TextMessage).Text("123"),
			}),
	}.Build()
)

func Test(t *testing.T) {
	var e = ExampleError()
	e.SetError(errors.New("Test error. "))

	if data, err := xml.MarshalIndent(e, "", "\t"); err != nil {
		t.Fatal(err)
	} else {
		fmt.Printf("%s\n", data)
	}
}