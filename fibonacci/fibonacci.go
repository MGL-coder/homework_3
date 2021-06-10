package fibonacci

func Fibonacci() func() int {
	f1 := 0
	f2 := 1

	return func() int {
		f := f1 + f2
		f1 = f2
		f2 = f
		return f1
	}
}
