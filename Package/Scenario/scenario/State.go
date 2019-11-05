package scenario

import "github.com/matscus/Hamster/Package/Hosts/hosts"

//State - struct for state scenario
type State struct {
	RunID      int64        `json:"runid"`
	Name       string       `json:"name"`
	Type       string       `json:"type"`
	StartTime  int64        `json:"starttime"`
	EndTime    int64        `json:"endtime"`
	Gun        string       `json:"gun"`
	Generators []hosts.Host `json:"generators"`
}
