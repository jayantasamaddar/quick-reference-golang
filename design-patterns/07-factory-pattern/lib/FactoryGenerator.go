package lib

type Employee struct {
	Name, Position, Designation string
	AnnualIncome                int
}

// Functional Approach
func NewEmployeeFactory(position string) func(name, designation string, annualIncome int) *Employee {
	return func(name, designation string, annualIncome int) *Employee {
		return &Employee{name, position, designation, annualIncome}
	}
}

// Structural Approach
type EmployeeFactory struct {
	Position string
}

func (f *EmployeeFactory) Create(name, designation string, annualIncome int) *Employee {
	return &Employee{name, f.Position, designation, annualIncome}
}

func NewEmployeeFactoryStructural(position string) *EmployeeFactory {
	return &EmployeeFactory{position}
}
