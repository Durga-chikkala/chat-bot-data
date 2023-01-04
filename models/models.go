package models

type QueriesData struct {
	Id       string `json:"id,omitempty"`
	Question string `json:"question,omitempty"`
	Solution string `json:"solution,omitempty"`
	Count    int64  `json:"count,omitempty"`
}
