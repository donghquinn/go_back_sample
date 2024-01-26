package calculator

import "errors"


func Sum(num1 int, num2 int) (int, error) {
	if (num1 < num2 ){
			return (num1 + num2) * (num2 - num1 + 1 ) / 2, nil
	} else {
		return -999999, errors.New("Please Provide Arguments Correctly")
	}
}