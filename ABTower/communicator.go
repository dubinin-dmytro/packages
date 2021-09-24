package ABTower

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Communicator struct {
	host  string
	token string
}

func (c Communicator) new(host string, token string) Communicator {
	return Communicator{
		host:  host,
		token: token,
	}
}

func (c *Communicator) call(url string) ([]byte, error) {
	fmt.Println("Call ", c.host+url)
	req, err := http.NewRequest("GET", c.host+url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", c.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
