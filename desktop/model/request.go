package model

type Request struct {
	ProjectId    uint64 `json:"project_id"`
	CollectionId uint64 `json:"collection_id"`
	Id           uint64 `json:"id"`
	Name         string `json:"name"`
	Type         string `json:"type"`
}

func (c Request) TableName() string {
	return TableRequest
}

func (c Request) GetProjectId() uint64 {
	return c.ProjectId
}

func (c Request) GetId() uint64 {
	return c.Id
}
