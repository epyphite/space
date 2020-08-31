package models

//FuelType for the tipe of fuel on rockets
type FuelType struct {
	ID          string
	Name        string
	Description string
}

//FuelCycle the type of changes between fuel used
type FuelCycle struct {
	ID          string
	Name        string
	Description string
}
