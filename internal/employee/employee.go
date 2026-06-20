package employee

// data pegawai
type Employee struct {
	ID 			int
	Name		string
	departement	string
	baseSalary	float64
}

// pegawai baru
func NewEmployee (id int, name, dept string, salary float64) Employee {
	return Employee{
	ID: id,
	Name: name,
	departement: dept,
	baseSalary: salary,
	}
}

func (e Employee) Departement() string {
	return e.departement
}