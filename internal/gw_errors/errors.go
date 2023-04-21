package gw_errors

import "errors"

const errUriNullMsg = "ErrUriNull"
const errEmptyRoutersMsg = "ErrEmptyRouters"

type ErrUriNull struct {
}

func (ErrUriNull) Error() error {
	return errors.New(errUriNullMsg)
}

type ErrEmptyRouters struct {
}

func (ErrEmptyRouters) Error() error {
	return errors.New(errEmptyRoutersMsg)
}
