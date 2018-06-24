package worldcup

import "strings"

const (
	RegularGoal = "goal"
	PenaltyGoal = "goal-penalty"
	OwnGoal     = "goal-own"

	SubstitutionIn  = "substitution-in"
	SubstitutionOut = "substitution-out"

	YellowCard = "yellow-card"
	RedCard    = "red-card"
)

// A TeamEvent represents a single team event resource received from http://worldcup.sfg.io/
type TeamEvent struct {
	ID          int    `json:"id"`
	Player      string `json:"player"`
	Time        string `json:"time"`
	TypeOfEvent string `json:"type_of_event"`
}

// IsGoal returns true if TeamEvent is a goal and false otherwise
func (event TeamEvent) IsGoal() bool {
	return strings.Contains(event.TypeOfEvent, "goal")
}

// IsPenalty returns true if TeamEvent is a penalty goal and false otherwise
func (event TeamEvent) IsPenalty() bool {
	return event.TypeOfEvent == PenaltyGoal
}

// IsOwnGoal returns true if TeamEvent is a own goal and false otherwise
func (event TeamEvent) IsOwnGoal() bool {
	return event.TypeOfEvent == OwnGoal
}

// IsSubstitution returns true if TeamEvent is a substitution and false otherwise
func (event TeamEvent) IsSubstitution() bool {
	return strings.Contains(event.TypeOfEvent, "substitution")
}

// IsSubstitutionIn returns true if TeamEvent is an in substitution and false otherwise
func (event TeamEvent) IsSubstitutionIn() bool {
	return event.TypeOfEvent == SubstitutionIn
}

// IsSubstitutionOut returns true if TeamEvent is an out substitution and false otherwise
func (event TeamEvent) IsSubstitutionOut() bool {
	return event.TypeOfEvent == SubstitutionOut
}

// IsCard returns true if TeamEvent is a card and false otherwise
func (event TeamEvent) IsCard() bool {
	return strings.Contains(event.TypeOfEvent, "card")
}

// IsYellowCard returns true if TeamEvent is a yellow card and false otherwise
func (event TeamEvent) IsYellowCard() bool {
	return event.TypeOfEvent == YellowCard
}

// IsRedCard returns true if TeamEvent is a red card and false otherwise
func (event TeamEvent) IsRedCard() bool {
	return event.TypeOfEvent == RedCard
}
