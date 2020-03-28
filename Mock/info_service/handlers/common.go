package handlers

type Balance struct {
	Type  string `json:"type"`
	Value struct {
		Amount   string `json:"amount"`
		Currency struct {
			NumCode  string `json:"numCode"`
			CharCode string `json:"charCode"`
		} `json:"currency"`
	} `json:"value"`
}
