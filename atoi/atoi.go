package atoi

import "errors"

func Atoi(s string) (int, error) {
	if len(s) == 0 {
		return 0, errors.New("empty string provided")
	}

	result := 0
	sign := 1

	i := 0
	if s[i] == '-' {
		sign = -1
		i++
	}

	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			result = result*10 + int(s[i]-'0')
		} else {
			return 0, errors.New("provided string is not a number")
		}
	}

	return sign * result, nil
}
