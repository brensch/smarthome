package smarthome

// Positions x,y
// 0,2 1,2 2,2
// 0,1 1,1 2,1
// 0,0 1,0 2,0

type ObjectType string

// type Appliance2 struct {
// 	Type     ApplianceType
// 	Location Location

// 	// -1 = going down
// 	// 0 = not set (default when viewing your own team)
// 	// 1 = going up
// 	Team int8

// 	Health   int8
// 	Strength int8

// 	// TODO
// 	Repair int8
// 	Model  int8

// 	// EventReceivers []EventReceiver
// }

type ObjectState struct {
	Location Location
	Upgrade  Upgrade

	// -1 = going down
	// 0 = not set (default when viewing your own team)
	// 1 = going up
	Team int8

	Health   int8
	Strength int8

	// TODO
	Repair int8
	Model  int8
}

type Location struct {
	X int8
	Y int8
}

var (
	TeamValues = []int8{
		-1,
		1,
	}
)

// team is either -1 or 1
// team 1 gets moved to top and inverted on both axis
func (l Location) MoveToStreet(width, height, team int8) Location {

	if team == 1 {
		return Location{
			X: l.X,
			Y: l.Y,
		}
	}

	// characters going down have been flipped
	return Location{
		X: width - l.X - 1,
		Y: height - l.Y - 1,
	}
}

// // this make any object that contains an object state implement the state method
// func (s ObjectState) State() ObjectState {
// 	return s
// }

// func (s ObjectState) SetLocation(loc Location) {
// 	s.Location = loc
// }

// this make any object that contains an object state implement the state method
func (s *ObjectState) State() ObjectState {
	return *s
}

func (s *ObjectState) SetLocation(loc Location) {
	s.Location = loc
}

// const (
// 	FriendsIndex = 0
// 	EnemiesIndex = 1
// )

type Appliance interface {
	State() ObjectState
	SetLocation(loc Location)

	// Type needs to just return the string of the type of the appliance
	Type() ObjectType

	// this allows the appliance to be converted to its street location (ie in the two player grid).
	// may decouple location somehow to remove this as a required method....
	MoveToStreet(width, height, team int8) Appliance

	// // CreateEvents performs all the logic based on an appliances current position and the positions
	// // of other appliances.
	// // The friends and enemies inputs are two separate Appliance slices so that you can pass in different arrays
	// // depending on the team of the appliance creating the events.
	// CreateEvents(appliances []Appliance) []Event

	// ReceiveEvent is how the appliance responds to events.
	// The returned appliance object is the next state of the appliance
	// All appliances receive all events in case they have behaviour based on things happening to
	// appliances elsewhere on the board.
	// Receiving events can also trigger more events, for example incoming damage to an ally might trigger a heal event etc
	ReceiveEvents(appliances []Appliance, events []Event, turn uint8) ([]Appliance, []Event)
}
