package version

import "fmt"

var (

	// URL is the git URL for the repository
	URL = "github.com/p9c/gel"
	// GitRef is the gitref, as in refs/heads/branchname
	GitRef = "refs/heads/master"
	// GitCommit is the commit hash of the current HEAD
	GitCommit = "dd4c2766a32b95de754b6768c0f9d8a70b7bcc19"
	// BuildTime stores the time when the current binary was built
	BuildTime = "2021-03-28T01:11:15+01:00"
	// Tag lists the Tag on the build, adding a + to the newest Tag if the commit is
	// not that commit
	Tag = "v0.1.0"
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
