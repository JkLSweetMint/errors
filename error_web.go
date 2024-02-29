package errors

// Web - абстракция системы web ошибок.
type Web interface {
	Error

	ToHttp() WebHttp
	ToWs() WebWs
}

// FieldsWeb - абстракция системы web ошибок с полями.
type FieldsWeb interface {
	Fields

	ToHttp() FieldsWebHttp
	ToWs() FieldsWebWs
}

// ----------------------------------------------------------------

// internalWeb - внутрення реализация web ошибки.
type internalWeb struct {
	internal
}

// ToHttp - преобразования ошибки в web http формат.
func (instance internalWeb) ToHttp() WebHttp {
	return WebHttp(internalWebHttp{
		internal{
			id:     instance.id,
			status: instance.status,
			t:      typeWebHttp,

			message: instance.message,
			err:     instance.err,

			httpStatusCode: instance.httpStatusCode,
		},
	})
}

// ToWs - преобразования ошибки в web ws формат.
func (instance internalWeb) ToWs() WebWs {
	return WebWs(internalWebWs{
		internal{
			id:     instance.id,
			status: instance.status,
			t:      typeWebWs,

			message: instance.message,
			err:     instance.err,

			wsStatusCode: instance.wsStatusCode,
		},
	})
}

// ----------------------------------------------------------------

// internalFieldsWeb - внутрення реализация web ошибок с полями.
type internalFieldsWeb struct {
	internalFields
}

// ToHttp - преобразования ошибки в web http формат.
func (instance internalFieldsWeb) ToHttp() FieldsWebHttp {
	return FieldsWebHttp(internalFieldsWebHttp{
		internalFields{
			internal: internal{
				id:     instance.id,
				status: instance.status,
				t:      typeWebHttp,

				message: instance.message,
				err:     instance.err,

				httpStatusCode: instance.httpStatusCode,
			},
			fields: instance.fields,
		},
	})
}

// ToWs - преобразования ошибки в web ws формат.
func (instance internalFieldsWeb) ToWs() FieldsWebWs {
	return FieldsWebWs(internalFieldsWebWs{
		internalFields{
			internal: internal{
				id:     instance.id,
				status: instance.status,
				t:      typeWebWs,

				message: instance.message,
				err:     instance.err,

				wsStatusCode: instance.wsStatusCode,
			},
			fields: instance.fields,
		},
	})
}
