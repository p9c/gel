package version

//go:generate go run ./update/.

import (
	"fmt"
)

var (

	// URL is the git URL for the repository
	URL = "github.com/p9c/gel"
	// GitRef is the gitref, as in refs/heads/branchname
	GitRef = "HEAD"
	// GitCommit is the commit hash of the current HEAD
	GitCommit = "e3bfce4ba69afe391a2979f922559eaa573f0b97"
	// BuildTime stores the time when the current binary was built
	BuildTime = "2021-04-30T18:12:27+02:00"
	// Tag lists the Tag on the build, adding a + to the newest Tag if the commit is
	// not that commit
	Tag = "v0.1.20"
	// PathBase is the path base returned from runtime caller
	PathBase = "/home/loki/src/github.com/p9c/pod/pkg/gel/"
	// Major is the major number from the tag
	Major = 0
	// Minor is the minor number from the tag
	Minor = 1
	// Patch is the patch version number from the tag
	Patch = 20
	// Meta is the extra arbitrary string field from Semver spec
	Meta = ""
)

// Get returns a pretty printed version information string
func Get() string {
	return fmt.Sprint(
		"\nRepository Information\n"+
		"\tGit repository: "+URL+"\n",
		"\tBranch: "+GitRef+"\n"+
		"\tCommit: "+GitCommit+"\n"+
		"\tBuilt: "+BuildTime+"\n"+
		"\tTag: "+Tag+"\n",
		"\tMajor:", Major, "\n",
		"\tMinor:", Minor, "\n",
		"\tPatch:", Patch, "\n",
		"\tMeta: ", Meta, "\n",
	)
}
