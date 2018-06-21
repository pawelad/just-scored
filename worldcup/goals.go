package worldcup

import "strings"

// A TeamEvent represents a single team event resource received from http://worldcup.sfg.io/
type TeamEvent struct {
	ID          int    `json:"id"`
	Player      string `json:"player"`
	Time        string `json:"time"`
	TypeOfEvent string `json:"type_of_event"`
}

// IsGoal returns true if passed TeamEvent is a goal and false otherwise
func (event TeamEvent) IsGoal() bool {
	return strings.Contains(event.TypeOfEvent, "goal")
}

// IsPenalty returns true if passed TeamEvent is a penalty goal and false otherwise
func (event TeamEvent) IsPenalty() bool {
	return event.TypeOfEvent == "goal-penalty"
}
