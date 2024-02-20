package sm_errors_http

import (
	errors "sm_errors"
	"sm_errors/entities"
)

var (
	ExampleBasicHttpError Basic = BasicConstructor{
		BasicConstructor: errors.BasicConstructor{
			ID:      0000000000000001,
			Status:  entities.StatusFailed,
			Message: new(entities.Message).Set("Example basic-http message. "),
		},
		StatusCode: StatusInternalServerError,
	}.Build()

	ExampleFieldsHttpError Fields = FieldsConstructor{
		FieldsConstructor: errors.FieldsConstructor{
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
