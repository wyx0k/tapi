package model

type Environment struct {
	ProjectId uint64 `json:"project_id"`
	Id        uint64 `json:"id"`
	Name      string `json:"name"`
}

func (e Environment) TableName() string {
	return "environment"
}

func (e Environment) GetProjectId() uint64 {
	return e.ProjectId
}

func (e Environment) GetId() uint64 {
	return e.Id
}
