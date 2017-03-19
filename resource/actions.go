package resource

import "github.com/hippoai/raappd/action"

type Actions []*action.Action

// NewActions instanciates
func NewActions() Actions {
	return []*action.Action{}
}

// Len
func (actions Actions) Len() int {
	return len(actions)
}

// Swap
func (actions Actions) Swap(i, j int) {
	actions[i], actions[j] = actions[j], actions[i]
}

// Less
func (actions Actions) Less(i, j int) bool {
	ai := actions[i].ExpectedPayloadDescription
	aj := actions[j].ExpectedPayloadDescription

	if ai.NRequired > aj.NRequired {
		return true
	}
	if ai.NRequired < aj.NRequired {
		return false
	}

	// Now if we are here they have the same number of required
	if ai.NOptional > aj.NOptional {
		return true
	}
	if ai.NOptional < aj.NOptional {
		return false
	}

	// If we land here this means they have the same numbers
	return false
}
