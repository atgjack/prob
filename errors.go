package distributions

import "errors"

// Signifies bad parameters for a distribution.
type InvalidParamsError struct{ S string }

// Signifies an undefined result. Eg. When the variance is undefined.
var IndeterminateError = errors.New("Iindeterminate value. Result is NaN.")

func (e InvalidParamsError) Error() string { return e.S }
