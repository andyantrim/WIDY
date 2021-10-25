package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

type CommitFilter struct {
	Author string
	Days   int
}

func (r Repo) filterCommits(c CommitFilter) (o []string, err error) {
	if r.Path == "" {
		return o, err
	}
	now := time.Now()
	if c.Days == 0 {
		c.Days = 1
		if now.Weekday() == time.Monday {
			c.Days = 3
		}
	}
	author := fmt.Sprintf("--author=%s", c.Author)
	y, m, d := time.Now().AddDate(0, 0, c.Days).Date()
	since := fmt.Sprintf("--since=%d/%d/%d", d, m, y)

	cmd := exec.Command("git", "-C", r.Path, "log", "-n", "3", author, "--pretty=format:%B", since, "--all")
	//fmt.Printf("Running: %s\n", cmd.String())

	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Failed to lookup repo %s\n", r.Path)
		return o, err
	}
	if len(outb.Bytes()) == 0 {
		return o, err
	}
	outString := outb.String()
	for _, s := range strings.Split(outString, "\n") {
		if s == "" {
			continue
		}
		o = append(o, fmt.Sprintf("%s: %s", r.Name, s))
	}
	return o, err
}
