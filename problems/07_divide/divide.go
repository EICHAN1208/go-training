package divide

type ErrZeroDivision struct{}

func (e ErrZeroDivision) Error() string {
	return "cannot divide by zero"
}

func Divide(dividend int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, ErrZeroDivision{}
	}
	return dividend / divisor, nil
}
