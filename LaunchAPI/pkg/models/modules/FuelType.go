package models

//FuelType for the tipe of fuel on rockets
type FuelType struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

//FuelCycle the type of changes between fuel used
type FuelCycle struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
