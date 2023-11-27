package variableouter

var (
	OuterVariable int
)

func SumTest(a, b int) int {
	OuterVariable = +(a + b)
	return a + b
}
