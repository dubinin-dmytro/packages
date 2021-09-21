package ABTower

type LapsResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    []Lap  `json:"data"`
}

type Lap struct {
	ID                    int         `json:"id"`
	ExperimentID          int         `json:"experiment_id"`
	RestartWithBranches   string      `json:"restart_with_branches"`
	PreservedBranchesUids interface{} `json:"preserved_branches_uids"`
	PreservedLapsIds      interface{} `json:"preserved_laps_ids"`
	NewUsersOnly          bool        `json:"new_users_only"`
	OldUsersOnly          bool        `json:"old_users_only"`
	PaidOrUnpaid          string      `json:"paid_or_unpaid"`
	UsersNumber           int         `json:"users_number"`
	CreatedAt             string      `json:"created_at"`
	UpdatedAt             string      `json:"updated_at"`
	StartedAt             string      `json:"started_at"`
	FinishedAt            string      `json:"finished_at"`
	DeletedAt             interface{} `json:"deleted_at"`
}

func (l LapsResponse) isSuccess() bool {
	return l.Success
}

func (l LapsResponse) errorMessage() string {
	if l.Success {
		return ""
	}
	return l.Message
}

func (l LapsResponse) getData() []Lap {
	if !l.Success {
		return []Lap{}
	}
	return l.Data
}
