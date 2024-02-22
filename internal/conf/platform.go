package conf

var (
	PF_GITHUB  = "Github"
	PF_GITEE   = "Gitee"
	PF_CURRENT = PF_GITHUB
)

func HasGithub() bool {
	return len(TNOTE_GIST_TOKEN) > 0
}

func IsGithub() bool {
	return PF_CURRENT == PF_GITHUB
}

func HasGitee() bool {
	return len(TNOTE_GIST_TOKEN_GITEE) > 0
}

func IsGitee() bool {
	return PF_CURRENT == PF_GITEE
}
