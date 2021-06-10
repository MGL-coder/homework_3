package itoa

func Itoa(n int) string {
	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	if n == 0 {
		return "0"
	}

	var str = ""
	for ; n > 0; n /= 10 {
		var digit = byte(n%10 + 48)
		str = string(digit) + str
	}

	return sign + str
}
