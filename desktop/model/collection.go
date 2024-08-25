package model

type CollectionTreeType string

const (
	CollectionType = "collection"
	RequestType    = "request"
)

type CollectionTree struct {
	ProjectId uint64             `json:"project_id"`
	Id        uint64             `json:"id"`
	ParentId  uint64             `json:"parent_id"`
	Name      string             `json:"name"`
	Type      CollectionTreeType `json:"type"`
}

func (c CollectionTree) TableName() string {
	return TableCollectionTree
}

func (c CollectionTree) GetProjectId() uint64 {
	return c.ProjectId
}

func (c CollectionTree) GetId() uint64 {
	return c.Id
}

type Collection struct {
	ProjectId uint64 `json:"project_id"`
	Id        uint64 `json:"id"`
}

func (c Collection) TableName() string {
	return TableCollection
}

func (c Collection) GetProjectId() uint64 {
	return c.ProjectId
}

func (c Collection) GetId() uint64 {
	return c.Id
}
