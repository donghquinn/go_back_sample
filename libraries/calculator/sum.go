package calculator

import "errors"

func Add(num1 float64, num2 float64) float64{
	return num1 + num2
}

func Sum(num1 int, num2 int) (int, error) {
	if (num1 < num2 ){
			return (num1 + num2) * (num2 - num1 ) / 2, nil
	} else {
		return -999999, errors.New("Please Provide Arguments Correctly")
	}
}