package common

import (
	"fmt"
	"goapp/internal/core/logic/common/errors"
)

func ErrorCode(err error) string {
	if kind, ok := err.(interface{ Kind() errors.Kind }); ok {
		return kind.Kind().String()
	}

	return ""
}

func IsClientSideError(err error) bool {
	if kind, ok := err.(interface{ Kind() errors.Kind }); ok {
		return kind.Kind() < errors.DevelopmentError
	}
	return false
}

func ValidationError(err error) error {
	return errors.Error(err, errors.ValidationError)
}

func UnsupportedQueryError(query Query) error {
	err := fmt.Errorf("unknown query type: %T", query)
	return errors.Error(err, errors.UnsupportedQueryError)
}
