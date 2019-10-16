package scenario

type TreadGroupsParams struct {
	TreadGroupsName  string             `json:"TreadGroupName"`
	TGType           string             `json:"TGType"`
	TreadGroupParams []TreadGroupParams `json:"TreadGroupParams"`
}
type TreadGroupParams struct {
	ParamType string `json:"ParamType"`
	Name      string `json:"Name"`
	Values    string `json:"Values"`
}
