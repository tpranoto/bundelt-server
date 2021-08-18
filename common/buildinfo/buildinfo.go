package buildinfo

import "fmt"

type buildInfo struct {
	gitCommit string
	branch    string
	buildNum  string
	buildDate string
}

func NewBuildInfo(gitCommit, branch, buildNum, buildDate string) buildInfo {
	var bNum string
	if buildNum == "" {
		bNum = "local"
	}

	return buildInfo{
		gitCommit: gitCommit,
		branch:    branch,
		buildNum:  bNum,
		buildDate: buildDate,
	}
}

func (b *buildInfo) String() string {
	format := `
	##### Build Info #####
	Commit hash:	%s
	Branch:			%s
	Build number:	%s
	Build time: 	%s
	`

	return fmt.Sprintf(format, b.gitCommit, b.branch, b.buildNum, b.buildDate)
}
