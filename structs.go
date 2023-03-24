package client

import "encoding/json"

// Machine represents response structure of machine.
type Machine struct {
	Hostname string            `yaml:"hostname" json:"hostname"`
	Active   bool              `yaml:"active" json:"active"`
	IPV4     string            `yaml:"ipv4" json:"ipv4"`
	IPV6     string            `yaml:"ipv6" json:"ipv6"`
	Labels   map[string]string `yaml:"labels" json:"labels"`
}

// String Returns string representation of machine.
func (m *Machine) String() string {
	bytes, _ := json.Marshal(m)
	return string(bytes)
}

// Status represents response if status is requested.
type Status struct {
	Auth     bool `json:"auth"`
	ReadOnly bool `json:"readonly"`
	Machines int  `json:"machines"`
}

// String Returns string representation of status.
func (s *Status) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}
