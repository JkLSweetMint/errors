package sm_errors

import "errors/entities"

var (
	ExampleBasicError Basic = ErrorConstructor{
		ID:      0000000000000001,
		Status:  entities.StatusFailed,
		Message: new(entities.Message).Set("Example basic message. "),
	}.Build()

	ExampleFieldsError FieldsError = FieldsErrorConstructor{
		ID:      0000000000000001,
		Status:  entities.StatusFailed,
		Message: new(entities.Message).Set("Example fields message. "),
		Fields: entities.Fields{
			{
				Key:     "login",
				Message: "Is required. ",
			},
		},
	}.Build()

	ExampleBasicErrorWithMessageFormat Basic = ErrorConstructor{
		ID:      0000000000000001,
		Status:  entities.StatusFailed,
		Message: new(entities.Message).Set("Example basic message with message formatting text={{text}}. "),
	}.Build()

	ExampleFieldsErrorWithMessageFormat FieldsError = FieldsErrorConstructor{
		ID:      0000000000000001,
		Status:  entities.StatusFailed,
		Message: new(entities.Message).Set("Example fields message with message formatting text={{text}}. "),
		Fields: entities.Fields{
			{
				Key:     "login",
				Message: "Is required. ",
			},
		},
	}.Build()
)
