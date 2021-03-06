/*
Copyright © 2021 Niklas Schlögel <niklasschloegel@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package tracking

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/niklasschloegel/parcly/config"
)

// Handles requests.
// If a body is specified through the responseStruct parameter,
// this method tries to marshal this struct into a json string.
// Responses are getting unmarshalled from a json string to a struct.
func DoRequest(method, url string, requestStruct, responseStruct interface{}) error {

	var request *http.Request
	var err error

	if requestStruct != nil {
		data, err := json.Marshal(requestStruct)
		if err != nil {
			return err
		}
		request, err = http.NewRequest(method, url, bytes.NewBuffer(data))
		if err != nil {
			return err
		}
	} else {
		request, err = http.NewRequest(method, url, nil)
		if err != nil {
			return err
		}
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Tracktry-Api-Key", config.TracktryApiKey)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		errorMsg := fmt.Sprintf("Error: %s", resp.Status)
		return errors.New(errorMsg)
	}

	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(responseBytes, &responseStruct); err != nil {
		anyResp := AnyResponse{}
		if err := json.Unmarshal(responseBytes, &anyResp); err != nil {
			return err
		}
		msg := fmt.Sprintf("%s [Code: %s]", anyResp.Meta.Message, strconv.Itoa(anyResp.Meta.Code))
		return errors.New(msg)
	}

	return nil
}
