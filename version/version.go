package version

import "fmt"

var (

	// URL is the git URL for the repository
	URL = "github.com/p9c/gel"
	// GitRef is the gitref, as in refs/heads/branchname
	GitRef = "refs/heads/master"
	// GitCommit is the commit hash of the current HEAD
	GitCommit = "6a195b1e804ea3e187c5bd4ad92c22320b859616"
	// BuildTime stores the time when the current binary was built
	BuildTime = "2021-03-28T01:36:56+01:00"
	// Tag lists the Tag on the build, adding a + to the newest Tag if the commit is
	// not that commit
	Tag = "v0.1.0+"
	// PathBase is the path base returned from runtime caller
	PathBase = "/home/loki/src/github.com/p9c/gel/"
)

// Get returns a pretty printed version information string
func Get() string {
	return fmt.Sprint(
		"ParallelCoin Pod\n"+
		"	git repository: "+URL+"\n",
		"	branch: "+GitRef+"\n"+
		"	commit: "+GitCommit+"\n"+
		"	built: "+BuildTime+"\n"+
		"	Tag: "+Tag+"\n",
	)
}
