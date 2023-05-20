package main

import "errors"

// Assuming we have 2 errors
func WrapErrors(errs ...error)(err error){
	err1 := errs[0]
	err2 := errs[1]

	return errors.Join(err1, err2)
}