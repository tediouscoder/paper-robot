package github

import "context"

// UpdateFile will update a file.
func UpdateFile(ctx context.Context, file string, fn func(string) (string, error)) (err error) {
	// Generate README.
	sha, originContent, err := GetFileContent(ctx, file)
	if err != nil {
		return
	}

	newContent, err := fn(originContent)
	if err != nil {
		return
	}

	return UpdateFileContent(ctx, file, sha, newContent)
}
