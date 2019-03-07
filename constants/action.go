package constants

// Action is the type used in paper robot action.
type Action uint8

// Actions that paper robot need to take.
const (
	ActionIgnore Action = iota
	ActionSetStateWaitReview
	ActionSetStateRequestChange
	ActionExecuteAdd
	ActionExecuteUpdate
	ActionExecuteRemove
	ActionClose
)
