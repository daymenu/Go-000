package errcode

import "fmt"

// Error error
type Error struct {
	Code Code
	Err  error
}

// Error error
func (err Error) Error() string {
	return fmt.Sprintf("code: %d, err: %v", err.Code, err.Err)
}

// UnWrap unwrap
func (err Error) UnWrap() error {
	return err.Err
}
