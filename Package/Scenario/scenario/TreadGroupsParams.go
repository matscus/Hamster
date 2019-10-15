package scenario

type TreadGroupsParams struct {
	TreadGroupsName  string `json:"TreadGroupsName"`
	TreadGroupParams []struct {
		ParamType string `json:"ParamType"`
		Name      string `json:"Name"`
		Values    string `json:"Values"`
	} `json:"TreadGroupParams"`
}
