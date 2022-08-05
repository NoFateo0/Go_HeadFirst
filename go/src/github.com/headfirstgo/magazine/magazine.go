package magazine

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	City   string
	Address
}

type Employee struct {
	Name   string
	Salary float64
	City   string
	Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}
