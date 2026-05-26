package model

type Damage struct {
	PedelecId string `json:"pedelecId"`
	Severity  int    `json:"severity"`
	Comment   string `json:"comment"`
}
