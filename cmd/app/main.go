package main

import (
	"fmt"

	"github.com/A-Fauzi-N/mini-project-1/internal/printer"

	"github.com/A-Fauzi-N/mini-project-1/internal/bonus"
	"github.com/A-Fauzi-N/mini-project-1/internal/employee"
)

const companyName = "Naf Corporation"

func main() {
	fmt.Println(
		"___ Welcome To",
		companyName,
		"Management System ___",
	)

	var fixedIDs [3]int = [3]int{
		101,
		102,
	}

	employees := []employee.Employee{
		employee.NewEmployee(
			fixedIDs[0],
			"Aldeb",
			"Engineering",
			6000000,
		),
		employee.NewEmployee(
			fixedIDs[1],
			"Rein",
			"HR",
			7000000,
		),
	}


	debtCount := make(map[string]int)

	for _, e := range employees {
		debtCount[e.Departement()]++
	}

	fmt.Println("\n>> Employee Distribution per Departement")

	for debt, count := range debtCount {
		fmt.Printf("- %s : %d\n", debt, count)
	}

	fmt.Println("\n>> Promoting Aldeb...")
	aldebPtr := &employees[0]
	aldebPtr.Promote(1000000)

	if emp, ok := employee.FindEmployee(101, employees); ok {
		fmt.Printf(
			"Found: %s with new salary: %.2f\n",
			emp.Name,
			emp.CalculatedPay(),
		)
	}

	totalPay := employee.TotalPayroll(
		employees[0],
		employees[1],
	)
	fmt.Printf(
		"\n>> Total Company Payroll: %.2f\n",
		totalPay,
	)

	fmt.Println("\n>> Calculating Bonus...")
	bonusCalc := bonus.CreateBonusCalculator(1.5)

	fmt.Printf(
		"Rein's Bonus: %.2f\n",
		 bonusCalc(employees[1].CalculatedPay()),
	)

	fmt.Println("\n>> Filtering Engineers...")
	engineers := employee.FilterEmployees(
		employees,
		func(e employee.Employee) bool {
			return e.Departement() == "Engineering"
		},
	)

	for _, eng := range engineers {
		fmt.Printf("- %s\n", eng.Name)
	}

	fmt.Println("\n>> Print Miscellanious Info:")

	printer.PrintInfo(
		"this demonstrates how any/interface{} works",
	)

	printer.PrintInfo(employees[1])

	printer.PrintInfo(404)
}
