package errors

import (
	"fmt"
)

type TemplateDAOError struct {
	ErrCode     uint16
	Description string
	Err         error
}

/*--------------------------Template DAO calls errors--------------------------*/

const (
	TemplateNotFoundInDB uint16 = 1 << iota
)

/*-------------------------------------------------------------------*/

var TemplateDAOErrorDescriptionMap = map[uint16]string{
	TemplateNotFoundInDB: "Template not found in db.",
}

func (r *TemplateDAOError) Error() string {
	return fmt.Sprintf("desc %v: err %v", r.Description, r.Err)
}

func NewTemplateDAOError(errCode uint16, err error) *TemplateDAOError {
	desc := TemplateDAOErrorDescriptionMap[errCode]

	return &TemplateDAOError{
		Description: desc,
		ErrCode:     errCode,
		Err:         err,
	}
}
