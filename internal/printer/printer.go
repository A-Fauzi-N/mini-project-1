package printer

import (
	"fmt"

	"github.com/A-Fauzi-N/mini-project-1/internal/employee"
)

func PrintInfo(data any) {
	switch v := data.(type) {
	case string:
		fmt.Println("[info - string]:", v)

	case employee.Employee:
		fmt.Printf(
			"[info - Employee]: %s (Debt: %s)\n",
			v.Name,
			v.Departement(),
		)
	case int:
		fmt.Println("[info - Numeric]:", v)

	default:
		fmt.Println("[info - Unknown]: Tipe data tidak dikenali")
	}
}