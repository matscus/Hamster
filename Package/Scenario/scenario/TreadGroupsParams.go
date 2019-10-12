package scenario

// type TreadGroupsParams struct {
// 	TreadGroupsName  string             `json:"treadgroupsname"`
// 	TreadGroupParams []TreadGroupParams `json:"treadgroupsparams"`
// }
// type TreadGroupParams struct {
// 	Name   string `json:"name"`
// 	Values string `json:"values"`
// }

type TreadGroupsParams struct {
	TreadGroupsName  string `json:"TreadGroupsName"`
	TreadGroupParams []struct {
		Name   string `json:"Name"`
		Values string `json:"Values"`
	} `json:"TreadGroupParams"`
}
