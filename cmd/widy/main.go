package main

import (
	"fmt"
	"os"

	"github.com/andyantrim/WIDY/git"
	"github.com/andyantrim/WIDY/walker"
)

func main() {
	w := walker.NewWalker("")
	contents, err := w.Walk()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := git.NewRepos(contents)
	results, err := r.GetCommitsByAuthor("andyantrim")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, r := range results {
		fmt.Println(r)
	}

}
