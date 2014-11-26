package types

type App struct {
	Id             string
	Name           string
	ReleaseID      string
	LastestVersion int
	CreateTime     int64
}

type DeployRequest struct {
	SvnURL      string `json:"svnURL,omitempty"`
	DockerImage string `json:"dockerImage,omitempty"`
}

type ScaleRequest struct {
	ProgramReplica map[string]int `json:"programReplica,omitempty"`
}
