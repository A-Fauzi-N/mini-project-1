package employee

func (e Employee) CalculatedPay() float64 {
	return e.baseSalary
}

func (e *Employee) Promote(raise float64) {
	e.baseSalary += raise
}