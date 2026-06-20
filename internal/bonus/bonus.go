package bonus

func CreateBonusCalculator(multiplier float64) func(float64) float64 {
	return func(salary float64) float64 {
		return salary * multiplier
	}
}