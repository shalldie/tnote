package conf

var (
	ENV_GITHUB  = "ENV_GITHUB"
	ENV_GITEE   = "ENV_GITEE"
	ENV_CURRENT = ENV_GITHUB
)

func IsGithub() bool {
	return ENV_CURRENT == ENV_GITHUB
}

func IsGitee() bool {
	return ENV_CURRENT == ENV_GITEE
}
