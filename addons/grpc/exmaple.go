package sm_errors_grpc

import (
	errors "errors"
	"errors/entities"
)

var (
	ExampleBasicGrpcError Error = ErrorConstructor{
		ErrorConstructor: errors.ErrorConstructor{
			ID:      0000000000000001,
			Status:  entities.StatusFailed,
			Message: new(entities.Message).Set("Example basic-grpc message. "),
		},
		StatusCode: StatusInternalServerError,
	}.Build()

	ExampleFieldsGrpcError FieldsError = FieldsErrorConstructor{
		FieldsErrorConstructor: errors.FieldsErrorConstructor{
			ID:      0000000000000001,
			Status:  entities.StatusFailed,
			Message: new(entities.Message).Set("Example fields-grpc message. "),
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
