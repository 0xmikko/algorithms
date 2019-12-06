package main

import (
	"fmt"
	"strconv"
)

func half(str string) (string, string) {
	n := len(str) / 2
	return str[:n/2], str[n/2:]

}


func alignNumbers(x, y string) (string, string) {
	if len(x) == len(y) {
		return x, y
	}

	if len(x) > len(y) {
		var newY string
		for i := 0; i < len(x) - len(y); i++ {
			newY += "0"
		}
		newY += y
		return x, y
	}

	y, x = alignNumbers(y, x)
	return x, y
}

func sum(x, y string) (string, error) {
	x, y = alignNumbers(x, y)
	sum := ""
	cum := 0
	for _, i := range x {
		a, err := strconv.Atoi(string(x[i]))
		if err != nil {
			return "", err
		}
		b, err := strconv.Atoi(string(y[i]))
		if err != nil {
			return "", err
		}

		sum += strconv.Itoa((a+b + cum) % 10)
		cum = (a + b + cum) / 10
	}

	return sum, nil
}

func multiply(x, y string) (string, error) {

	if len(x) == 1 && len(y)== 1 {

		aInt, err := strconv.Atoi(x)
		if err != nil {
			return "", err
		}


		bInt, err := strconv.Atoi(y)
		if err != nil {
			return "", err
		}
		return fmt.Sprint(aInt * bInt), nil
	}

	a, b := half(x)
	c, d := half(y)

	ac, err := multiply(a, c)
	if err != nil {
		return "", err
	}

	bd, err := multiply(b, d)
	if err != nil {
		return "", err
	}

	ad, err := multiply(a, d)
	if err != nil {
		return "", err
	}

	bc, err := multiply(b, c)
	if err != nil {
		return "", err
	}







	return ac + bd + ad + bc, nil

}

func main(){

}
