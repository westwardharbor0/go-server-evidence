package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// ServerEvidenceClient represents necessary structure to communicate with service.
type ServerEvidenceClient struct {
	Host string
}

func (sec *ServerEvidenceClient) doRequest(
	method, endpoint string, payload io.Reader,
) (*http.Response, error) {
	apiURL, err := url.JoinPath(sec.Host, endpoint)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, apiURL, payload)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	if err != nil {
		return nil, err
	}

	return client.Do(request)
}

// All returns every machine in service.
func (sec *ServerEvidenceClient) All() ([]Machine, error) {
	var machines []Machine
	resp, err := sec.doRequest("GET", "/machines/", nil)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.Unmarshal(bodyBytes, &machines); err != nil {
		return nil, err
	}

	return machines, nil
}

// Filter returns subset of machines based on supplied fields.
func (sec *ServerEvidenceClient) Filter(
	field, value string,
) ([]Machine, error) {
	var machines []Machine
	endpoint := fmt.Sprintf("/machines/%s/%s", field, value)
	resp, err := sec.doRequest("GET", endpoint, nil)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bodyBytes, &machines); err != nil {
		return nil, err
	}

	return machines, nil
}

// Status returns the stats about running service.
func (sec *ServerEvidenceClient) Status() (*Status, error) {
	var status Status
	resp, err := sec.doRequest("GET", "/status", nil)

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := json.Unmarshal(bodyBytes, &status); err != nil {
		return nil, err
	}

	return &status, nil
}

// Update updates the machines attributes in service.
// Update only works if machines in service are not read only.
// Check Status() for this info.
func (sec *ServerEvidenceClient) Update(machine Machine) error {
	jsonBytes, err := json.Marshal(machine)
	if err != nil {
		return err
	}

	resp, err := sec.doRequest("PUT", "/machines/", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf(
			"service returned %d status code on machine update",
			resp.StatusCode,
		)
	}

	return nil
}

// Delete removes machine from service.
// Delete only works if machines in service are not read only.
// Check Status() for this info.
func (sec *ServerEvidenceClient) Delete(machine Machine) error {
	endpoint := fmt.Sprintf("/machines/%s", machine.Hostname)
	resp, err := sec.doRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		return fmt.Errorf(
			"service returned %d status code on machine delete",
			resp.StatusCode,
		)
	}

	return nil
}
