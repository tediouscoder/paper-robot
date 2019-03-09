package paper

import (
	"context"

	"gopkg.in/go-playground/webhooks.v5/github"

	"github.com/tediouscoder/paper-robot/constants"
	ig "github.com/tediouscoder/paper-robot/internal/github"
	"github.com/tediouscoder/paper-robot/model"
	"github.com/tediouscoder/paper-robot/utils"
)

// Handler will handle the github issue event.
func Handler(ctx context.Context, payload *github.IssuesPayload) (err error) {
	action := utils.CalculateAction(payload)
	switch action {
	case constants.ActionSetStateWaitReview:
		err = setStateWaitReview(ctx)
	case constants.ActionSetStateRequestChange:
		err = setStateRequestChange(ctx)
	case constants.ActionExecuteAdd:
		err = executeAdd(ctx)
	case constants.ActionExecuteUpdate:
		err = executeUpdate(ctx)
	case constants.ActionExecuteRemove:
		err = executeRemove(ctx)
	case constants.ActionClose:
		err = closeIssue(ctx)
	default:
		// Not implement action and ignore action will do nothing.
		return
	}

	if err == nil {
		return
	}

	switch err {
	case constants.ErrRequiredFiledMissing:
		return
	case constants.ErrInvalidPaperURL:
		return
	default:
		return
	}
}

func setStateWaitReview(ctx context.Context) (err error) {
	err = ig.CreateIssueComment(ctx, "Please waiting for review.")
	if err != nil {
		return
	}

	return ig.AddIssueLabel(ctx, string(constants.StateWaitReview))
}

func setStateRequestChange(ctx context.Context) (err error) {
	// Remove wait review before set wait change.
	err = ig.RemoveIssueLabel(ctx, string(constants.StateWaitReview))
	if err != nil {
		return
	}

	return ig.AddIssueLabel(ctx, string(constants.StateWaitChange))
}

func executeAdd(ctx context.Context) (err error) {
	// Remove wait review before execute.
	err = ig.RemoveIssueLabel(ctx, string(constants.StateWaitReview))
	if err != nil {
		return
	}

	// Parse issue
	paper, err := model.ParsePaper(ctx)
	if err != nil {
		return
	}

	// Update data.json
	sha, dataContent, err := ig.GetFileContent(ctx, constants.DataFilePath)
	if err != nil {
		return
	}
	data, err := model.ParseData(dataContent)
	if err != nil {
		return
	}

	err = data.AddPaper(paper)
	if err != nil {
		return
	}

	dataContent, err = model.FormatData(data)
	if err != nil {
		return
	}
	err = ig.UpdateFileContent(ctx, constants.DataFilePath, sha, dataContent)
	if err != nil {
		return
	}

	// Generate README.
	sha, _, err = ig.GetFileContent(ctx, constants.ReadmeFilePath)
	if err != nil {
		return
	}

	content, err := GenerateREADME(data)
	if err != nil {
		return
	}

	err = ig.UpdateFileContent(ctx, constants.ReadmeFilePath, sha, content)
	if err != nil {
		return
	}

	err = ig.AddIssueLabel(ctx, string(constants.StateFinished))
	if err != nil {
		return
	}

	err = ig.CloseIssue(ctx)
	if err != nil {
		return
	}
	return
}

func executeUpdate(ctx context.Context) (err error) {
	return
}

func executeRemove(ctx context.Context) (err error) {
	return
}

func closeIssue(ctx context.Context) (err error) {
	return ig.CloseIssue(ctx)
}
