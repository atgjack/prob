package distributions

import "errors"

// Signifies an undefined result. Eg. When the variance is undefined.
type InvalidParamsError struct{ S string }

// Signifies an undefined result. Eg. When the variance is undefined.
var IndeterminateError = errors.New("Iindeterminate value.")

func (e InvalidParamsError) Error() string { return e.S }
