package model

type Hook interface {
	GetProjectId() string
	SetProjectId(string)
	GetUrl() string
}

type GithubHookConf struct {
	ContentType string `json:"content_type"`
	Url         string `json:"url"`
	InsecureSsl string `json:"insecure_ssl"`
}

//Todo : one project may have multiple hooks related to our system, but now we restrict only one

type GithubHook struct {
	ProjectId string         `json:"project_id" bson:"_id"`
	Id        int            `json:"id"`
	Events    []string       `json:"events"`
	Config    GithubHookConf `json:"config"`
	Type      string         `json:"type"`
	Name      string         `json:"name"`
	Active    bool           `json:"active"`
	Url       string         `json:"url"`
	TestUrl   string         `json:"test_url"`
	PingUrl   string         `json:"ping_url"`
}

func (gh *GithubHook) GetProjectId() string {
	return gh.ProjectId
}
func (gh *GithubHook) SetProjectId(projectId string) {
	gh.ProjectId = projectId
}
func (gh *GithubHook) GetUrl() string {
	return gh.Url
}
