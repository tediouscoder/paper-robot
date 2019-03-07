package constants

import "strings"

// Label is type paper robot need to use.
type Label string

// Request that paper robot supported.
const (
	RequestAdd    Label = "request/add"
	RequestUpdate Label = "request/update"
	RequestRemove Label = "request/remove"
)

// Review operations that paper robot provided.
const (
	ReviewApprove       Label = "review/approve"
	ReviewReject        Label = "review/reject"
	ReviewRequestChange Label = "review/request_change"
)

// State that paper robot recognized.
const (
	StateWaitReview Label = "state/wait_review"
	StateWaitChange Label = "state/wait_change"
	StateFailed     Label = "state/failed"
	StateFinished   Label = "state/finished"
)

// Type's that paper robot will use.
const (
	TypeRequest = "request"
	TypeReview  = "review"
	TypeState   = "state"
)

// Type will get label's type.
func (l Label) Type() string {
	switch v := strings.Split(string(l), "/")[0]; v {
	case TypeRequest, TypeReview, TypeState:
		return v
	default:
		return ""
	}
}
