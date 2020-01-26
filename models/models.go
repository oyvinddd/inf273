package models

func NewVehicle(index int, home int, start int, cap int) *Vehicle {
	return &Vehicle{index: index, home: home, start: start, cap: cap}
}

type Vehicle struct {
	index int
	home  int
	start int
	cap   int
}
