package models

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
	return &Vehicle{}
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

// Vehicle struct
type Vehicle struct {
	Index    int
	Home     int
	Start    int
	Capacity int
}

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
