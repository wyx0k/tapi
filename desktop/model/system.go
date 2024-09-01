package model

type Preference struct {
	Common PreferencesCommon `json:"common"`
	Data   PreferencesData   `json:"data"`
	Hub    PreferencesHub    `json:"hub"`
}

type PreferencesCommon struct {
	Theme    string `json:"theme"`
	Language string `json:"language"`
}

type PreferencesData struct {
	DataDir string `json:"data_dir"`
}
type PreferencesHub struct {
	HubAddress string `json:"hub_address"`
}

type AppUIState struct {
	CurrentProjectId uint64 `json:"current_project_id,omitempty"`
	CollectionTabs   []Tab  `json:"collection_tabs"`
	EnvironmentTabs  []Tab  `json:"environment_tabs"`
	CurrentSessionId uint64 `json:"current_session_id,omitempty"`
}
