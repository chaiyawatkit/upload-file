package errors

// ParameterError Invalid parameter
type ParameterError struct {
	Message string `json:"message"`
}

func (err ParameterError) Error() string {
	return err.Message
}

// UnprocessableEntity Valid parameter but invalid business and etc.
type UnprocessableEntity struct {
	Message string `json:"message"`
}

func (err UnprocessableEntity) Error() string {
	return err.Message
}

// RecordNotFoundError Cannot find resource.
type RecordNotFoundError struct {
	Message string `json:"message"`
}

func (err RecordNotFoundError) Error() string {
	return err.Message
}

// ForeignKeyConstraintError Invalid parameter
type ForeignKeyConstraintError struct {
	Message string `json:"message"`
}

func (err ForeignKeyConstraintError) Error() string {
	return err.Message
}

// UniqueKeyConstraintError Invalid parameter
type UniqueKeyConstraintError struct {
	Message string `json:"message"`
}

func (err UniqueKeyConstraintError) Error() string {
	return err.Message
}

// InternalError Database error and etc.
type InternalError struct {
	Message string `json:"message"`
}

func (err InternalError) Error() string {
	return err.Message
}
