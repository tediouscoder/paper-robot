package github

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/google/go-github/github"

	"github.com/tediouscoder/paper-robot/constants"
	"github.com/tediouscoder/paper-robot/utils"
)

// AddIssueLabel will add a label for current issue.
func AddIssueLabel(ctx context.Context, label string) (err error) {
	em := utils.FromEventMetadataContext(ctx)

	_, _, err = client.Issues.AddLabelsToIssue(ctx, em.Owner, em.Repo, int(em.IssueNumber), []string{label})
	if err != nil {
		return
	}
	return
}

// RemoveIssueLabel will remove a label from current issue.
func RemoveIssueLabel(ctx context.Context, label string) (err error) {
	em := utils.FromEventMetadataContext(ctx)

	_, err = client.Issues.RemoveLabelForIssue(ctx, em.Owner, em.Repo, int(em.IssueNumber), label)
	if err != nil {
		return
	}
	return
}

// GetFileContent will get a file's content.
func GetFileContent(ctx context.Context, file string) (sha, data string, err error) {
	em := utils.FromEventMetadataContext(ctx)

	content, _, _, err := client.Repositories.GetContents(ctx, em.Owner, em.Repo, file, nil)
	if err != nil {
		return
	}

	sha = content.GetSHA()

	data, err = content.GetContent()
	if err != nil {
		return
	}
	return
}

// UpdateFileContent will update a file's content.
func UpdateFileContent(ctx context.Context, file, sha, content string) (err error) {
	em := utils.FromEventMetadataContext(ctx)

	message := fmt.Sprintf("[paper-robot] Update file %s via issue #%d", filepath.Base(file), em.IssueNumber)

	_, _, err = client.Repositories.UpdateFile(
		ctx, em.Owner, em.Repo, file,
		&github.RepositoryContentFileOptions{
			Message: &message,
			Content: []byte(content),
			SHA:     &sha,
			Committer: &github.CommitAuthor{
				Name:  github.String(constants.CommitAuthorName),
				Email: github.String(constants.CommitAuthorEmail),
			},
		})
	if err != nil {
		return
	}
	return
}

// CloseIssue will close a issue.
func CloseIssue(ctx context.Context) (err error) {
	em := utils.FromEventMetadataContext(ctx)

	_, _, err = client.Issues.Edit(ctx, em.Owner, em.Repo, int(em.IssueNumber), &github.IssueRequest{
		State: github.String("closed"),
	})
	if err != nil {
		return
	}
	return
}

// GetIssueContent will get an issue's content.
func GetIssueContent(ctx context.Context) (content string, err error) {
	em := utils.FromEventMetadataContext(ctx)

	issue, _, err := client.Issues.Get(ctx, em.Owner, em.Repo, int(em.IssueNumber))
	if err != nil {
		return
	}

	return issue.GetBody(), nil
}
