package distributions

// Signifies bad parameters for a distribution.
type InvalidParamsError struct{ S string }

func (e InvalidParamsError) Error() string { return e.S }
