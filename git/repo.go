package git

import (
	"strings"
)

type Repo struct {
	Path string
	Name string
}

type Repos []Repo

func NewRepos(paths []string) Repos {
	r := make(Repos, len(paths))
	for _, p := range paths {
		parts := strings.Split(p, "/")
		name := parts[len(parts)-1]
		r = append(r, Repo{
			Path: p,
			Name: name,
		})
	}

	return r
}

func (r Repo) GetCommitsByAuthor(author string, days int) ([]string, error) {
	f := CommitFilter{
		Author: author,
		Days: days,
	}
	return r.filterCommits(f)
}

func (r Repos) GetCommitsByAuthor(author string, days int) (c []string, err error) {
	var s []string
	for _, r1 := range r {
		s, err = r1.GetCommitsByAuthor(author, days)
		if len(s) > 0 {
			c = append(c, s...)
		}
	}

	return c, err
}
