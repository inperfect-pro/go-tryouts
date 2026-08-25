package main

import "fmt"

type Payable interface {
	fmt.Stringer
	CalculatePay() float64
}

type SalariedEmployee struct {
	Name         string
	AnnualSalary float64
}

func (se SalariedEmployee) CalculatePay() float64 {
	return se.AnnualSalary / 12.0
}

func (se SalariedEmployee) String() string {
	return fmt.Sprintf("Salaried: %s (Annual: $%.2f)", se.Name, se.AnnualSalary)
}

type HourlyEmployee struct {
	Name        string
	HourlyRate  float64
	HoursWorked float64
}

func (he HourlyEmployee) CalculatePay() float64 {
	return he.HourlyRate * he.HoursWorked
}

func (he HourlyEmployee) String() string {
	return fmt.Sprintf("Hourly: %s (Rate: $%.2f/hr, Hours: %.1f)", he.Name, he.HourlyRate, he.HoursWorked)
}

type CommissionEmployee struct {
	Name           string
	BaseSalary     float64
	CommissionRate float64
	SalesAmount    float64
}

func (ce CommissionEmployee) CalculatePay() float64 {
	return ce.BaseSalary + (ce.CommissionRate * ce.SalesAmount)
}

func (ce CommissionEmployee) String() string {
	return fmt.Sprintf("Commission: %s (Base: $%.f Rate: $%.2f/hr, Sales: %.1f)", ce.Name, ce.BaseSalary, ce.CommissionRate, ce.SalesAmount)
}

func PrintEmployeeSummary[P fmt.Stringer](employee P) {
	fmt.Printf("  - Processing: %s\n", employee)
}

func ProcessPayroll(employees []Payable) {
	fmt.Println("\n--- Processing Payroll ---")
	totalPayroll := 0.0
	for _, employee := range employees {
		PrintEmployeeSummary(employee)
		pay := employee.CalculatePay()
		fmt.Printf("    Monthly Pay: $%.2f\n", pay)
		totalPayroll += pay
	}
	fmt.Printf("\nTotal Monthly Payroll: $%.2f\n", totalPayroll)
	fmt.Println("\n--------------------------")
}

func main() {
	fmt.Println("Welcome to the Payroll Processor!")

	salEmp := SalariedEmployee{"Jane", 33000}
	hrEmp := HourlyEmployee{
		Name:        "Jon",
		HourlyRate:  12.43,
		HoursWorked: 34,
	}
	commEmp := CommissionEmployee{
		Name:           "Jess",
		BaseSalary:     1000,
		CommissionRate: 0.1,
		SalesAmount:    3567,
	}

	payrollList := []Payable{
		salEmp,
		hrEmp,
		commEmp,
		HourlyEmployee{"Diana", 30.00, 150.0},
	}

	ProcessPayroll(payrollList)

}
