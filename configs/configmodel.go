package configs

type Configurations struct {
	Github   GithubConfiguration
	Crucible CrucibleConfiguration
}

type Access struct {
	User    string
	Key     string
	BaseUrl string
}

type RepositoryOptions struct {
	Repository string
	Branch     string
	Owner      string
}

type GithubConfiguration struct {
	Access       Access
	Repositories []RepositoryOptions
}

type CrucibleProjectOptions struct {
	Name string
}

type CrucibleConfiguration struct {
	Access   Access
	Projects []CrucibleProjectOptions
}
