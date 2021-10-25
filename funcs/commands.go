package funcs

import (
	"fmt"
	"path/filepath"

	"github.com/andyantrim/WIDY/git"
	"github.com/andyantrim/WIDY/walker"
	"github.com/urfave/cli/v2"
)

func runner(c *cli.Context) error {
	author := c.String("author")
	days := c.Int("days")

	file := c.String("dir")
	path, err := filepath.Abs(file)
	if err != nil {
		return fmt.Errorf("failed to find abs path for %s", file)
	}

	w := walker.NewWalker(path)
	contents, err := w.Walk()
	if err != nil {
		return err
	}

	r := git.NewRepos(contents)
	results, _ := r.GetCommitsByAuthor(author, days)
	if len(results) == 0 {
		fmt.Printf("Found no commits in %s\n", path)
	}
	for _, r := range results {
		fmt.Println(r)
	}

	return nil
}
