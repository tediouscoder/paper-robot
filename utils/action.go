package utils

import (
	"gopkg.in/go-playground/webhooks.v5/github"

	"github.com/tediouscoder/paper-robot/constants"
)

// CalculateAction will calculate actions via current issues pay load.
func CalculateAction(payload *github.IssuesPayload) constants.Action {
	// We should always ignore issue that not labeled.
	if payload.Action != constants.GithubIssueEventActionLabeled {
		return constants.ActionIgnore
	}

	// In labeled event, this must not be nil, but we should handle it.
	if payload.Label == nil || len(payload.Issue.Labels) == 0 {
		return constants.ActionIgnore
	}

	// Get current label.
	l := constants.Label((*payload.Label).Name)

	// Get current action.
	var action constants.Label
	for _, v := range payload.Issue.Labels {
		l := constants.Label(v.Name)
		if l.Type() != constants.TypeRequest {
			continue
		}

		action = l
		break
	}

	switch l.Type() {
	case constants.TypeRequest:
		if len(payload.Issue.Labels) > 1 {
			// Add a request label in an issue with more than one label
			// is strange, and we should ignore it for safe.
			return constants.ActionIgnore
		}

		switch l {
		case constants.RequestAdd, constants.RequestRemove, constants.RequestUpdate:
			// We need to set wait review state is new request come in.
			return constants.ActionSetStateWaitReview
		default:
			return constants.ActionIgnore
		}

	case constants.TypeReview:
		switch l {
		case constants.ReviewReject:
			// If review results in reject, we should close the issue.
			return constants.ActionClose

		case constants.ReviewRequestChange:
			// If review results in request change, we should update the state.
			return constants.ActionSetStateRequestChange

		case constants.ReviewApprove:
			// If review results in approve, we should execute the related action.
			switch action {
			case constants.RequestAdd:
				return constants.ActionExecuteAdd
			case constants.RequestUpdate:
				return constants.ActionExecuteUpdate
			case constants.RequestRemove:
				return constants.ActionExecuteRemove
			default:
				return constants.ActionIgnore
			}

		default:
			return constants.ActionIgnore
		}

	case constants.TypeState:
		// State label is used to mark current issue's state, paper robot
		// should ignore it.
		return constants.ActionIgnore

	default:
		// If label's type is not the one that paper robot will handle,
		// we should ignore it.
		return constants.ActionIgnore
	}
}
