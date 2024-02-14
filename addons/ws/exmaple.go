package sm_errors_ws

import (
	errors "errors"
	"errors/entities"
)

var (
	ExampleBasicWsError Error = ErrorConstructor{
		ErrorConstructor: errors.ErrorConstructor{
			ID:      0000000000000001,
			Status:  entities.StatusFailed,
			Message: new(entities.Message).Set("Example basic-ws message. "),
		},
		StatusCode: StatusInternalServerError,
	}.Build()

	ExampleFieldsWsError FieldsError = FieldsErrorConstructor{
		FieldsErrorConstructor: errors.FieldsErrorConstructor{
			ID:      0000000000000001,
			Status:  entities.StatusFailed,
			Message: new(entities.Message).Set("Example fields-ws message. "),
			Fields: entities.Fields{
				{
					Key:     "login",
					Message: "Is required. ",
				},
			},
		},
		StatusCode: StatusInternalServerError,
	}.Build()
)
