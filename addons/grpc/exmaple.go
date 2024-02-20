package sm_errors_grpc

import (
	errors "sm_errors"
	"sm_errors/entities"
)

var (
	ExampleBasicGrpcError Basic = BasicConstructor{
		BasicConstructor: errors.BasicConstructor{
			ID:      0000000000000001,
			Status:  entities.StatusFailed,
			Message: new(entities.Message).Set("Example basic-grpc message. "),
		},
		StatusCode: StatusInternalServerError,
	}.Build()

	ExampleFieldsGrpcError Fields = FieldsConstructor{
		FieldsConstructor: errors.FieldsConstructor{
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
