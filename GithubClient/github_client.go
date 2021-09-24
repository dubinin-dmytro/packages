package GithubClient

import (
	"github.com/google/go-github/github"
)

type GithubClient interface {
	New(oauthToken string) GithubClient
	GetPullRequestInfoFull(info PRInfo) (PRResponseInfo, error)
	GetPullRequestInfoFullMultiple(infos map[int]PRInfo) (map[int]*PRResponseInfo, error)
}

type PRResponseInfo struct {
	Id          int
	Info        PRInfo
	PullRequest github.PullRequest
	Reviews     []*github.PullRequestReview
	Error       []error
}

func GetGithubClient(token string, mode string) GithubClient {

	switch mode {
	case "ASYNC":
		return AsyncGithubClient{}.New(token)
	case "SYNC":
		return SyncGithubClient{}.New(token)
	default:
		return SyncGithubClient{}.New(token)
	}

	return nil
}
