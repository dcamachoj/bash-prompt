package main

func newFrm(values ...int) []int {
	return values
}

func formatSet(value int) []int {
	return newFrm(value)
}

func formatRst(value int) []int {
	if value == 0 {
		return newFrm(0)
	}
	return newFrm(20 + value)
}

func formatFg(value int) []int {
	return newFrm(30 + value)
}

func formatBg(value int) []int {
	return newFrm(40 + value)
}

func formatFg256(value int) []int {
	return newFrm(38, 5, value)
}

func formatBg256(value int) []int {
	return newFrm(48, 5, value)
}
