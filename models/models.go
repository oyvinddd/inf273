package models

import "fmt"

// --------------- VEHICLE ---------------

// Vehicle struct
type Vehicle struct {
	Index    int
	Home     int
	Start    int
	Capacity int
}

// NewVehicle convenience constructor for vehicle struct
func NewVehicle(index int, home int, start int, capacity int) *Vehicle {
	return &Vehicle{
		Index:    index,
		Home:     home,
		Start:    start,
		Capacity: capacity,
	}
}

// NewDummyVehicle convenience constructor for dummy vehicle
func NewDummyVehicle() *Vehicle {
	return &Vehicle{
		Index: -1,
	}
}

// IsDummy returns true if instance is a dummy vehicle else false
func (v Vehicle) IsDummy() bool {
	return v.Index < 0
}

func (v Vehicle) String() string {
	return fmt.Sprintf("[ vehicle | %-5v | %-5v | %5v | %-5v ]", v.Index, v.Home, v.Start, v.Capacity)
}

// --------------- CALL ---------------

// Call struct
type Call struct {
	Index       int
	Origin      int
	Destination int
	Size        int
	Penalty     int
	LowerPW     int
	UpperPW     int
	LowerDW     int
	UpperDW     int
}

// NewCall convenience constructor for call struct
func NewCall(index int, origin int, destination int, size int, penalty int, lpw int, upw int, ldw int, udw int) *Call {
	return &Call{
		Index:       index,
		Origin:      origin,
		Destination: destination,
		Size:        size,
		Penalty:     penalty,
		LowerPW:     lpw,
		UpperPW:     upw,
		LowerDW:     ldw,
		UpperDW:     udw,
	}
}

func (c Call) String() string {
	return fmt.Sprintf("[ call | %-5v | %-5v | %-5v | %-5v | %-5v | %-5v | %-5v | %-5v ]", c.Index, c.Origin, c.Destination, c.Size, c.LowerPW, c.UpperPW, c.LowerDW, c.UpperDW)
}

// --------------- TIME AND COST ---------------

// TimeAndCost struct
type TimeAndCost struct {
	Time int
	Cost int
}

// --------------- DATA ---------------

// INF273Data is a container for all data from file
type INF273Data struct {
	Vehicles []*Vehicle
	Calls    []*Call
}
