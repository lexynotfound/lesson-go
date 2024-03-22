package main

import "fmt"

// Vechicle adalah struktur dasar untuk kendaraan

// FuelType adalah tipe data untuk jenis bahan bakar
type FuelType string

const (
	Gasoline  FuelType = "Gasoline"
	Pertamax  FuelType = "Pertamax"
	Pertalite FuelType = "Pertalite"
	Battrey   FuelType = "Battrey"
)

type Vehicle struct {
	Name      string
	FuelLevel int
	Fuel      FuelType
}

func (v *Vehicle) RefillFuel(amount int) {
	v.FuelLevel += amount
	fmt.Printf("%s telah disi ulang dengan %d unit bahan bakar %s. Level bahan bakar saat ini: %d\n", v.Name, amount, v.Fuel, v.FuelLevel)
}

// Car
type Car struct {
	Vehicle
	Make string
}

//Motorcycle

type Motorcycle struct {
	Vehicle
	Make string
}

func NewCar(name, make string, fuel FuelType) *Car {
	return &Car{
		Vehicle: Vehicle{
			Name:      name,
			FuelLevel: 0,
			Fuel:      fuel,
		},
		Make: make,
	}
}

// NewMotorcycle membuat instance baru dari Motorcycle
func NewMotorcycle(name, make string, fuel FuelType) *Motorcycle {
	return &Motorcycle{
		Vehicle: Vehicle{
			Name:      name,
			FuelLevel: 0,
			Fuel:      fuel,
		},
		Make: make,
	}
}

func main() {
	hondaCar := NewCar("Honda Civic", "Honda", Gasoline)
	teslaCar := NewCar("Tesla Model S", "Tesla", Gasoline)
	yamahaMotorcycle := NewMotorcycle("Yamaha R6", "Yamaha", Gasoline)

	hondaCar.RefillFuel(20)
	teslaCar.RefillFuel(30)
	yamahaMotorcycle.RefillFuel(10)

	// Contoh untuk mengisi dengan Pertamax dan Pertalite
	hondaCar.Fuel = Pertamax
	hondaCar.RefillFuel(10)

	yamahaMotorcycle.Fuel = Pertalite
	yamahaMotorcycle.RefillFuel(5)
}
