package sm_errors_http

import (
	errors "errors"
	"errors/entities"
)

var (
	ExampleBasicHttpError Error = ErrorConstructor{
		ErrorConstructor: errors.ErrorConstructor{
			ID:      0000000000000001,
			Status:  entities.StatusFailed,
			Message: new(entities.Message).Set("Example basic-http message. "),
		},
		StatusCode: StatusInternalServerError,
	}.Build()

	ExampleFieldsHttpError FieldsError = FieldsErrorConstructor{
		FieldsErrorConstructor: errors.FieldsErrorConstructor{
			ID:      0000000000000001,
			Status:  entities.StatusFailed,
			Message: new(entities.Message).Set("Example fields-http message. "),
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
