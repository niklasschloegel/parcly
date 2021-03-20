package tracking

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/niklasschloegel/parcly/config"
)

func DoRequest(method, url string, requestStruct, responseStruct interface{}) error {

	var request *http.Request
	var err error

	if requestStruct != nil {
		data, err := json.Marshal(requestStruct)
		if err != nil {
			return err
		}
		request, err = http.NewRequest(method, url, bytes.NewBuffer(data))
	} else {
		request, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Tracktry-Api-Key", config.TracktryApiKey)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(responseBytes, &responseStruct); err != nil {
		return err
	}

	return nil
}
