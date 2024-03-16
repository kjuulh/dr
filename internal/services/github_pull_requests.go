package services

import "github.com/google/uuid"

type GitHubPullRequest struct {
	Title        string
	Description  string
	Comments     []string
	StatusChecks []string
	Diff         string
}

func newBogusPr() GitHubPullRequest {
	uuid := uuid.NewString()

	return GitHubPullRequest{
		Title:       "some pr" + uuid,
		Description: "some long text" + uuid,
		Comments: []string{
			"some comment" + uuid,
			"some comment" + uuid,
		},
		StatusChecks: []string{
			"some status check" + uuid,
			"some status check" + uuid,
		},
		Diff: "some diff" + uuid,
	}
}

func newBogusPrs(amount int) []GitHubPullRequest {
	prs := make([]GitHubPullRequest, 0, amount)

	for range amount {
		prs = append(prs, newBogusPr())
	}

	return prs
}

type GitHubPullRequestService struct {
	prs []GitHubPullRequest
}

func NewGitHubPullRequestService() *GitHubPullRequestService {
	return &GitHubPullRequestService{
		prs: newBogusPrs(50),
	}
}

func (g *GitHubPullRequestService) GetNext() (pr *GitHubPullRequest, ok bool) {
	if len(g.prs) > 0 {
		pr := g.prs[0]
		g.prs = g.prs[1:]
		return &pr, true
	}

	return nil, false
}
