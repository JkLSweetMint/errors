package errors

import (
	"reflect"
	"sm-errors/entities/details"
	"sm-errors/entities/messages"
	"sm-errors/types"
	"testing"
)

func TestConstructor_Build_WithError(t *testing.T) {
	type testCase[T Error] struct {
		name string
		c    Constructor[T]
		want T
	}

	tests := []testCase[Error]{
		{
			name: "Case 1",
			c: Constructor[Error]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
			},
			want: ExampleError(),
		},
		{
			name: "Case 2",
			c: Constructor[Error]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),
			},
			want: ExampleErrorWithDetails(),
		},
		{
			name: "Case 3",
			c: Constructor[Error]{
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
			},
			want: ExampleErrorWithDetailsAndFields(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Build()(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor_fillEmptyField_WithError(t *testing.T) {
	type testCase[T Error] struct {
		name string
		c    Constructor[T]
		want *Constructor[T]
	}

	tests := []testCase[Error]{
		{
			name: "Case 1",
			c: Constructor[Error]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
			},
			want: &Constructor[Error]{
				ID:     "T-000001",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error. "),
				Details: new(details.Details),
			},
		},
		{
			name: "Case 2",
			c: Constructor[Error]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),
			},
			want: &Constructor[Error]{
				ID:     "T-000002",
				Type:   types.TypeSystem,
				Status: types.StatusFatal,

				Message: new(messages.TextMessage).
					Text("Example error with details. "),
				Details: new(details.Details).
					Set("key", "value"),
			},
		},
		{
			name: "Case 3",
			c: Constructor[Error]{
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
			},
			want: &Constructor[Error]{
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
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.fillEmptyField(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fillEmptyField() = %v, want %v", got, tt.want)
			}
		})
	}
}
