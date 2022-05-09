package main

import (
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func main() {

	auth := http.TokenAuth{Token: "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"}
	opts := git.CloneOptions{
		URL:  "https://github.com/Alan-Jino/Address",
		Auth: &auth,
	}
	git.PlainClone("/tmp/cloneDir", false, &opts)

}
