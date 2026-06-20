package employee

import "github.com/A-Fauzi-N/mini-project-1/internal/interfaces"

func FindEmployee(id int, emps []Employee) (Employee, bool) {
	for _, e := range emps {
		if e.ID == id {
			return e, true
		}
	}
	return Employee{}, false
}

func TotalPayroll(payables ...interfaces.Payable) float64 {
	var total float64
	for _, p := range payables {
		total += p.CalculatedPay()
	}
	return total
}

func FilterEmployees(
	emps []Employee,
	filterFunc func(Employee) bool,
) []Employee {
	var filtered []Employee

	for _, e := range emps {
		if filterFunc(e) {
			filtered = append(filtered, e)
		}
	}
	return filtered
}
