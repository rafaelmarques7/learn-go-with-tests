package iterations

const repeatCount = 5

func Repeat(s string, repeatTimes int) (repeated string) {
	// default to 5 times, if an invalid argument is passed
	if repeatTimes <= 0 {
		repeatTimes = repeatCount
	}

	for i := 0; i < repeatTimes; i++ {
		repeated += s
	}
	return repeated
}
