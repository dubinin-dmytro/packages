package ABTower

type ExperimentsResponse struct {
	Success bool         `json:"success"`
	Message string       `json:"message"`
	Data    []Experiment `json:"data"`
}

type Experiment struct {
	ID                     int         `json:"id"`
	ProjectID              int         `json:"project_id"`
	UID                    string      `json:"uid"`
	ReportID               interface{} `json:"report_id"`
	CategoryID             int         `json:"category_id"`
	PlaceID                interface{} `json:"place_id"`
	Token                  interface{} `json:"token"`
	Name                   string      `json:"name"`
	Description            interface{} `json:"description"`
	Type                   string      `json:"type"`
	Issue                  string      `json:"issue"`
	WinnerStatus           interface{} `json:"winner_status"`
	WinnerBranchID         interface{} `json:"winner_branch_id"`
	CountEventsAfterFinish int         `json:"count_events_after_finish"`
	SaveWayGetIntoExp      bool        `json:"save_way_get_into_exp"`
	SaveMicrotime          bool        `json:"save_microtime"`
	SaveUserExperiments    bool        `json:"save_user_experiments"`
	CreatedAt              string      `json:"created_at"`
	UpdatedAt              string      `json:"updated_at"`
	FinishedAt             interface{} `json:"finished_at"`
	ArchivedAt             interface{} `json:"archived_at"`
	HoldingFrom            interface{} `json:"holding_from"`
	DeletedAt              interface{} `json:"deleted_at"`
}

func (exp ExperimentsResponse) isSuccess() bool {
	return exp.Success
}

func (exp ExperimentsResponse) errorMessage() string {
	if exp.Success {
		return ""
	}
	return exp.Message
}

func (exp ExperimentsResponse) getData() []Experiment {
	if !exp.Success {
		return []Experiment{}
	}
	return exp.Data
}
