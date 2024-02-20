package sm_errors_ws

import (
	errors "sm_errors"
	"sm_errors/entities"
)

var (
	ExampleBasicWsError Basic = BasicConstructor{
		BasicConstructor: errors.BasicConstructor{
			ID:      0000000000000001,
			Status:  entities.StatusFailed,
			Message: new(entities.Message).Set("Example basic-ws message. "),
		},
		StatusCode: StatusInternalServerError,
	}.Build()

	ExampleFieldsWsError Fields = FieldsConstructor{
		FieldsConstructor: errors.FieldsConstructor{
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
