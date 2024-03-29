package model

// Request ...
type RequestTask struct {
	A float64 `json:"A"`
	B float64 `json:"B"`
	C float64 `json:"C"`
}

// Response ...
type ResponseCalculation struct {
	Name   string `json:"Name"`
	Status string `json:"Status"`
	Result string `json:"Result"`
	Host   string
	Port   string
}
