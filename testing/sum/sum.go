package sum

func Sum(ints... int) int  {
	if len(ints) == 0 {
		return 0
	}
	return ints[0] + Sum(ints[1:]...)
}
