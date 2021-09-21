package ABTower

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

const DateLayout = "2006-01-02"

type Processor struct {
	communicator Communicator
}

func (p Processor) New(host string, token string) Processor {
	return Processor{
		Communicator{}.new(host, token),
	}
}

func (p *Processor) operateCall(url string, response ResponseInterface) error {
	res, err := p.communicator.call(url)

	if err != nil {
		return err
	}

	err = json.Unmarshal(res, &response)

	if err != nil {
		return err
	}

	if !response.isSuccess() {
		return errors.New(response.errorMessage())
	}

	return nil
}

func (p *Processor) GetLapInfoByPeriod(lapId int, df time.Time, dto time.Time) (LapInfo, error) {
	url := fmt.Sprintf("/api/analytics/?lapId=%d&periodFrom=%s&periodTo=%s",
		lapId,
		df.Format(DateLayout),
		dto.Format(DateLayout),
	)

	resp := LapInfoResponse{}
	err := p.operateCall(url, &resp)
	return resp.getData(), err
}

func (p *Processor) GetLaps(expId int) ([]Lap, error) {
	url := fmt.Sprintf("/api/experiments/%d/laps", expId)

	resp := LapsResponse{}
	err := p.operateCall(url, &resp)
	return resp.getData(), err
}

func (p *Processor) GetAllExperiments(categoryUid string) ([]Experiment, error) {
	url := fmt.Sprintf("/api/experiments/?filter[category_uids]=%s", categoryUid)

	resp := ExperimentsResponse{}
	err := p.operateCall(url, &resp)
	return resp.getData(), err
}

func (p *Processor) GetActiveExperiments(categoryUid string) ([]Experiment, error) {
	url := fmt.Sprintf("/api/experiments/?filter[category_uids]=%s&filter[status]=active", categoryUid)

	resp := ExperimentsResponse{}
	err := p.operateCall(url, &resp)
	return resp.getData(), err
}
