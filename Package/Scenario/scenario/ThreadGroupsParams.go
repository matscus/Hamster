package scenario

//ThreadGroup - struct for the jmeter scenario tread groups params
type ThreadGroup struct {
	ThreadGroupName   string              `json:"threadGroupName"`
	ThreadGroupType   string              `json:"threadGroupType"`
	ThreadGroupParams []ThreadGroupParams `json:"threadGroupParams"`
}

//ThreadGroupParams - simple jmeter tread  groups param from TreadGroupsParams struct
type ThreadGroupParams struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
