package lib

/*********************************************************************************************************/
// Builder Facets: PersonBuilder with PersonAddressBuilder and PersonWorkBuilder
/*********************************************************************************************************/
type Person struct {
	Name string

	// Address Information
	Street, City, Zip string

	// Work Information
	Company, Designation string
	AnnualIncome         int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder(name string) *PersonBuilder {
	return &PersonBuilder{&Person{Name: name}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}
type PersonWorkBuilder struct {
	PersonBuilder
}

// Utility Methods to access the PersonAddressBuilder and PersonWorkBuilder
func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*pb}
}
func (pb *PersonBuilder) Works() *PersonWorkBuilder {
	return &PersonWorkBuilder{*pb}
}

// PersonAddressBuilder specific methods
func (pb *PersonAddressBuilder) At(street string) *PersonAddressBuilder {
	pb.person.Street = street
	return pb
}
func (pb *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	pb.person.City = city
	return pb
}
func (pb *PersonAddressBuilder) Zip(zip string) *PersonAddressBuilder {
	pb.person.Zip = zip
	return pb
}

// PersonWorkBuilder specific methods
func (pb *PersonWorkBuilder) At(company string) *PersonWorkBuilder {
	pb.person.Company = company
	return pb
}
func (pb *PersonWorkBuilder) Designation(designation string) *PersonWorkBuilder {
	pb.person.Designation = designation
	return pb
}
func (pb *PersonWorkBuilder) AnnualIncome(income int) *PersonWorkBuilder {
	pb.person.AnnualIncome = income
	return pb
}

// Build the actual Person
func (pb *PersonBuilder) Build() *Person {
	return pb.person
}
