package model

type Project struct {
	Id    uint64 `json:"id"`
	HubId uint64 `json:"hub_id"`
	Name  string `json:"name"`
}

func (c Project) TableName() string {
	return TableProject
}

func (c Project) GetProjectId() uint64 {
	return c.Id
}

func (c Project) GetId() uint64 {
	return c.Id
}
