// This package wraps the API for requesting Enka.Network
package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MingxuanGame/enkanetwork-go/model"
)

// GetInfo requests Enka.Network and returns the EnkaNetworkData structure, or nil if there is an error
func GetInfo(uid uint32) (data *model.EnkaNetworkData, err error) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", fmt.Sprintf("https://enka.network/u/%d/__data.json", uid), nil)
	request.Header.Add("User-Agent", "EnkaNetwork-Go/0.1.0")
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http: status \"%s\" is not \"200 OK\"", resp.Status)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var enkaData model.EnkaNetworkData
	if err := json.Unmarshal(body, &enkaData); err != nil {
		return nil, err
	}
	return &enkaData, nil
}
