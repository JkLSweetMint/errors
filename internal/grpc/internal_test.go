package grpc

import (
	"reflect"
	"sm-errors/entities/details"
	"sm-errors/entities/messages"
	"sm-errors/internal"
	"sm-errors/types"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		store *internal.Store
	}

	tests := []struct {
		name  string
		args  args
		wantI *Internal
	}{
		{
			name: "Case 1",
			args: args{
				store: &internal.Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error. "),
				},
			},
			wantI: &Internal{
				Internal: internal.New(&internal.Store{
					ID:     "T-000001",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error. "),
				}),
			},
		},
		{
			name: "Case 2",
			args: args{
				store: &internal.Store{
					ID:     "T-000002",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error with details. "),
					Details: new(details.Details).
						Set("key", "value"),
				},
			},
			wantI: &Internal{
				Internal: internal.New(&internal.Store{
					ID:     "T-000002",
					Type:   types.TypeSystem,
					Status: types.StatusFatal,

					Message: new(messages.TextMessage).
						Text("Example error with details. "),
					Details: new(details.Details).
						Set("key", "value"),
				}),
			},
		},
		{
			name: "Case 3",
			args: args{
				store: &internal.Store{
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
			wantI: &Internal{
				Internal: internal.New(&internal.Store{
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
				}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := New(tt.args.store); !reflect.DeepEqual(gotI, tt.wantI) {
				t.Errorf("New() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}
