package parser

type Test struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Parent    bool    `json:"parent"`
	Children  []*Test `json:"children"`
	Result    string  `json:"result"`
	Duration  float64 `json:"duration"`
	Systemout string  `json:"stdout"`
}
